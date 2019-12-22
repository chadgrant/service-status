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

delimiter ;

drop procedure if exists deployment_paged;

delimiter |
create procedure deployment_paged(
    in p_environment varchar(150),
    in p_service varchar(150),
    in p_page int,
    in p_size int,
    out p_total int)
begin
    declare offset int;
    set offset = (p_page * p_size) - p_size;
    set offset = if(offset<0,0,offset);

    set p_environment = if(p_environment = '',null,p_environment);
    set p_service = if(p_service = '',null,p_service);

    select count(*) into p_total
    from deployment d
    inner join environment e on d.environment_id = e.id
    inner join service s on d.service_id = s.id
    where s.friendly_name = coalesce(p_service, s.friendly_name)
      and e.friendly_name = coalesce(p_environment, e.friendly_name);

    select d.id, e.friendly_name as environment, s.friendly_name as service, d.build_number, d.created, d.updated
    from deployment d
    inner join environment e on d.environment_id = e.id
    inner join service s on d.service_id = s.id
    where s.friendly_name = coalesce(p_service, s.friendly_name)
      and e.friendly_name = coalesce(p_environment, e.friendly_name)
    order by d.created desc, d.build_number desc
    limit p_size offset offset;
end|

delimiter ;
