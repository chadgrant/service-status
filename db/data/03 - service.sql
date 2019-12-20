use service_status;

drop procedure if exists service_upsert;

delimiter |
create procedure service_upsert(
    in p_friendly varchar(150),
    in p_name varchar(150),
    in p_friendly_name varchar(150),
    in p_status varchar(150),
    in p_active bit,
    in p_sort tinyint)
begin
    if nullif(trim(p_friendly),'') is null or not exists(select id from service where friendly_name=trim(p_friendly)) then
        insert into service (name,friendly_name,status,active,sort)
        values(trim(p_name),trim(p_friendly_name),p_status,p_active,p_sort);
    else
        update service
        set name=trim(p_name), friendly_name=trim(p_friendly_name), status=p_status, active=p_active, sort=p_sort, updated=current_timestamp()
        where friendly_name=trim(p_friendly);
    end if;
end|

delimiter ;

drop procedure if exists service_delete;

delimiter |
create procedure service_delete(in p_friendly varchar(150))
begin
    delete from service where friendly_name=trim(p_friendly);
end|