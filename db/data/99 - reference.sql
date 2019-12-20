use service_status;

call environment_upsert(null, 'Development', 'development', 1, 40);
call environment_upsert(null, 'QA', 'qa', 1, 30);
call environment_upsert(null, 'Staging', 'staging', 1, 20);
call environment_upsert(null, 'Production', 'production', 1, 10);

call service_upsert(null, 'Sample Service', 'sample_service', 'pending', 1, 40);

call deployment_insert('development', 'sample_service', '1.0.0', @dev_deploy);
call deployment_insert('qa', 'sample_service', '1.0.0', @qa_deploy);
call deployment_insert('staging', 'sample_service', '1.0.0', @stg_deploy);
call deployment_insert('production', 'sample_service', '1.0.0', @prd_deploy);

call service_environment_upsert('development', 'sample_service', @dev_deploy);
call service_environment_upsert('qa', 'sample_service', @qa_deploy);
call service_environment_upsert('staging', 'sample_service', @stg_deploy);
call service_environment_upsert('production', 'sample_service', @prd_deploy);

call service_instance_upsert('', 'development', 'sample_service', 'hostname-123', @dev_deploy, 'pending', 'http://dev1-sample.localdomain.com');
call service_instance_upsert('', 'qa', 'sample_service', 'hostname-245', @qa_deploy, 'pending', 'http://qa1-sample.localdomain.com');
call service_instance_upsert('', 'staging', 'sample_service', 'hostname-666', @stg_deploy, 'pending', 'http://stg1-sample.localdomain.com');
call service_instance_upsert('', 'production', 'sample_service', 'hostname-999', @prd_deploy, 'pending', 'http://prd1-sample.localdomain.com');

call service_endpoint_insert('development', 'sample_service', 'http://sample.local.com', @dev_svc_ep);
call service_endpoint_insert('qa', 'sample_service', 'http://sample.qa.com', @qa_svc_ep);
call service_endpoint_insert('staging', 'sample_service', 'http://sample.staging.com', @stg_svc_ep);
call service_endpoint_insert('production', 'sample_service', 'http://sample.production.com', @prd_svc_ep);