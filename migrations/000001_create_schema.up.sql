CREATE TABLE IF NOT EXISTS projects (
    id SERIAL,
    name character varying(500) NOT NULL,
    description character varying(1000),
    CONSTRAINT projects_pkey PRIMARY KEY (id),
    CONSTRAINT name_not_empty CHECK (name::text <> ''::text)
);

CREATE TABLE IF NOT EXISTS columns (
    id SERIAL,
    name character varying(255) DEFAULT 'ToDo'::character varying,
    "position" integer DEFAULT 1,
    project_id integer NOT NULL,
    CONSTRAINT columns_pkey PRIMARY KEY (id),
    CONSTRAINT name_unique UNIQUE (name, project_id),
    CONSTRAINT columns_project_id_fkey FOREIGN KEY (project_id)
        REFERENCES projects (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT name_not_empty CHECK (name::text <> ''::text)
);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL,
    name character varying(500) NOT NULL,
    description character varying(5000),
    "position" integer NOT NULL,
    column_id integer NOT NULL,
    CONSTRAINT tasks_pkey PRIMARY KEY (id),
    CONSTRAINT tasks_column_id_fkey FOREIGN KEY (column_id)
        REFERENCES columns (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT name_not_empty CHECK (name::text <> ''::text)
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL,
    text character varying(5000) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    task_id integer NOT NULL,
    CONSTRAINT comments_pkey PRIMARY KEY (id),
    CONSTRAINT comments_task_id_fkey FOREIGN KEY (task_id)
        REFERENCES tasks (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT text_not_empty CHECK (text::text <> ''::text)
);

CREATE OR REPLACE FUNCTION delete_column(
    src_cid integer)
    RETURNS integer
    LANGUAGE 'plpgsql'
AS $BODY$
declare
    dst_pos int;
    src_pid int;
    src_pos int;
    dst_cid int;
    tsk_pos int;
    tsk record;
    cln record;
begin
    select project_id into src_pid from columns where id=src_cid;
    select count(id) into dst_pos from columns where project_id=src_pid;
    if dst_pos>1 then
        select position into src_pos from columns where id=src_cid;
        if src_pos>1 then
            dst_pos=src_pos-1;
        end if;
        select id into dst_cid from columns where position=dst_pos;
        select count(id)+1 into tsk_pos from tasks where column_id=dst_cid;
        for tsk in select * from tasks where column_id=src_cid order by position
        loop
            update tasks set position=tsk_pos, column_id=dst_cid where id=tsk.id;
            tsk_pos = tsk_pos + 1;
        end loop;
        delete from columns where id=src_cid;
        update columns set position=position-1 where project_id=src_pid and position>src_pos;
        return dst_cid;
    end if;
    return -1 as dst_cid;
end;
$BODY$;