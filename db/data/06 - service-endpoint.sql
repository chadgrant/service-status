use service_status;

drop procedure if exists service_endpoint_insert;

delimiter |
create procedure service_endpoint_insert(
    in p_environment varchar(150),
    in p_service varchar(150),
    in p_endpoint varchar(512),
    out service_endpoint_id bigint)
begin
    select id into @env from environment where friendly_name=p_environment;
    select id into @svc from service where friendly_name=p_service;

    insert into service_endpoint(environment_id, service_id, endpoint)
    values(@env,@svc,p_endpoint);

    set service_endpoint_id = last_insert_id();
end|

delimiter ;

drop procedure if exists service_endpoint_delete;

delimiter |
create procedure service_endpoint_delete(
    in p_environment varchar(150),
    in p_service varchar(150),
    in p_endpoint varchar(512))
begin
    select id into @env from environment where friendly_name=p_environment;
    select id into @svc from service where friendly_name=p_service;

    delete from service_endpoint 
    where environment_id=@env
      and service_id=@svc
      and endpoint=p_endpoint;
end|