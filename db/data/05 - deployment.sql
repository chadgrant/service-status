use service_status;

drop procedure if exists deployment_insert;

delimiter |
create procedure deployment_insert(
    in p_environment varchar(150),
    in p_service varchar(150),
    in p_build_number varchar(150),
    out p_deployment_id bigint)
begin
    select id into @env from environment where friendly_name=p_environment;
    select id into @svc from service where friendly_name=p_service;

    insert into deployment(environment_id, service_id, build_number)
    values(@env,@svc,p_build_number);

    set p_deployment_id = last_insert_id();
end|

delimiter ;

drop procedure if exists deployment_delete;

delimiter |
create procedure deployment_delete(
    in p_environment varchar(150),
    in p_service varchar(150),
    in p_build_number varchar(150))
begin
    select id into @env from environment where friendly_name=p_environment;
    select id into @svc from service where friendly_name=p_service;

    delete from service_endpoint 
    where environment_id=@env
      and service_id=@svc
      and build_number=p_build_number;
end|