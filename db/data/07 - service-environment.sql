use service_status;

drop procedure if exists service_environment_upsert;

delimiter |
create procedure service_environment_upsert(
    in p_environment varchar(150),
    in p_service varchar(150),
    in p_deployment_id bigint)
begin
    select id into @env from environment where friendly_name=p_environment;
    select id into @svc from service where friendly_name=p_service;

    if not exists(select id from service_environment where environment_id=@env and service_id=@svc) then
        insert ignore into service_environment (environment_id, service_id, deployment_id)
        values(@env, @svc, p_deployment_id);
    else
        update service_environment
           set deployment_id=p_deployment_id, updated=current_timestamp()
        where environment_id=@env 
          and service_id=@svc;
    end if;
end|

delimiter ;

drop procedure if exists service_environment_delete;

delimiter |
create procedure service_environment_delete(
    in p_environment varchar(150),
    in p_service varchar(150))
begin
    select id into @env from environment where friendly_name=p_environment;
    select id into @svc from service where friendly_name=p_service;

    delete from service_environment 
    where environment_id=@env
      and service_id=@svc;
end|