use service_status;

drop procedure if exists service_instance_upsert;

delimiter |
create procedure service_instance_upsert(
    in p_name varchar(150),
    in p_environment varchar(150),
    in p_service varchar(150),
    in p_newname varchar(150),
    in p_deployment_id bigint,
    in p_status varchar(150),
    in p_endpoint varchar(512))
begin
    select id into @env from environment where friendly_name=p_environment;
    select id into @svc from service where friendly_name=p_service;

    if nullif(trim(p_name),'') is null or not exists(select id from service_instance where environment_id=@env and service_id=@svc and name=p_name) then
        insert into service_instance (environment_id, service_id, deployment_id, status, name, endpoint)
        values(@env, @svc, p_deployment_id, p_newname, p_status, p_endpoint);
    else
        update service_instance
           set name=p_newname, status=p_status, endpoint=p_endpoint, updated=current_timestamp()
        where environment_id=@env
          and service_id=@svc
          and name=p_name;
    end if;
end|

delimiter ;

drop procedure if exists service_instance_delete;

delimiter |
create procedure service_instance_delete(
    in p_environment varchar(150),
    in p_service varchar(150),
    in p_name varchar(150))
begin
    select id into @env from environment where friendly_name=p_environment;
    select id into @svc from service where friendly_name=p_service;

    delete from service_instance
    where environment_id=@env
      and service_id=@svc
      and name=@p_name;
end|