use service_status;

drop procedure if exists environment_upsert;

delimiter |
create procedure environment_upsert(
    in p_friendly varchar(150),
    in p_name varchar(150),
    in p_friendly_name varchar(150),
    in p_active bit,
    in p_sort tinyint)
begin
    if nullif(trim(p_friendly),'') is null or not exists(select id from environment where friendly_name=trim(p_friendly)) then
        insert into environment (name,friendly_name,active,sort)
        values(trim(p_name),trim(p_friendly_name),p_active,p_sort);
    else
        update environment 
        set name=trim(p_name), friendly_name=trim(p_friendly_name), active=p_active, sort=p_sort, updated=current_timestamp()
        where friendly_name=trim(p_friendly);
    end if;
end|

delimiter ;

drop procedure if exists environment_delete;

delimiter |
create procedure environment_delete(in p_friendly varchar(150))
begin
    delete from environment where friendly_name=trim(p_friendly);
end|