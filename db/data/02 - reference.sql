use service_status;

insert ignore into environment (name)
    values('Development'),
          ('QA'),
          ('Staging'),
          ('Production');

select id into @development_env from environment where name='Development';
select id into @qa_env from environment where name='QA';
select id into @staging_env from environment where name='Staging';
select id into @production_env from environment where name='Production';

insert ignore into service (name,friendly_name)
    values('sample_service','Sample Service');

select id into @sample_svc from service where name='sample_service';

insert ignore into service_endpoint (environment_id, service_id, endpoint)
    values (@development_env, @sample_svc, 'http://sample.localdomain.com'),
           (@qa_env, @sample_svc, 'http://sample.qa.com'),
           (@staging_env, @sample_svc, 'http://sample.staging.com'),
           (@production_env, @sample_svc, 'http://sample.production.com');

insert ignore into deployment (environment_id, service_id, build_number)
    values(@development_env, @sample_svc, '1.0.0'),
           (@qa_env, @sample_svc, '1.0.0'),
           (@staging_env, @sample_svc, '1.0.0'),
           (@production_env, @sample_svc, '1.0.0');

select id into @development_deploy from deployment where environment_id=@development_env and service_id=@sample_svc order by created desc limit 1;
select id into @qa_deploy from deployment where environment_id=@qa_env and service_id=@sample_svc order by created desc limit 1;
select id into @staging_deploy from deployment where environment_id=@staging_env and service_id=@sample_svc order by created desc limit 1;
select id into @production_deploy from deployment where environment_id=@production_env and service_id=@sample_svc order by created desc limit 1;

insert ignore into service_environment (environment_id, service_id, deployment_id)
    values(@development_env, @sample_svc, @development_deploy),
           (@qa_env, @sample_svc, @qa_deploy),
           (@staging_env, @sample_svc, @staging_deploy),
           (@production_env, @sample_svc, @production_deploy);

insert ignore into service_instance (environment_id, service_id, deployment_id, name, endpoint)
    values(@development_env, @sample_svc, @development_deploy, 'hostname-123', 'http://dev1-sample.localdomain.com'),
          (@qa_env, @sample_svc, @qa_deploy, 'hostname-245','http://qa1-sample.localdomain.com'),
          (@staging_env, @sample_svc, @staging_deploy, 'hostname-678', 'http://staging1-sample.localdomain.com'),
          (@production_env, @sample_svc, @production_deploy, 'hostname-993', 'http://prod1-sample.localdomain.com');