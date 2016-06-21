#include <string>

#include <grpc/grpc.h>
#include <grpc++/server.h>
#include <grpc++/server_builder.h>
#include <grpc++/server_context.h>
#include <grpc++/security/server_credentials.h>
#include "driverservice.grpc.pb.h"
#include "driver.grpc.pb.h"
#include "l2service.grpc.pb.h"
#include "port.h"

extern "C" {
#include "sal/driver.h"
#include "sal/version.h"
}

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::ServerReader;
using grpc::ServerReaderWriter;
using grpc::ServerWriter;
using grpc::Status;
using l2service::InitRequest;
using l2service::InitResponse;
using l2service::L2;

class DriverServiceImpl final : public driverservice::Driver::Service {
    public:
        Status Init(ServerContext* context, const driver::InitRequest* req, driver::InitResponse* res) {
            opennsl_init_t config;
            int ret = 0;
            ret = opennsl_driver_init(&config);
            if( ret == 0 ){
                return Status::OK;
            }
            return Status(grpc::UNAVAILABLE, "");
        }
        Status GetVersion(ServerContext* context, const driver::GetVersionRequest* req, driver::GetVersionResponse* res) {
            res->set_version(opennsl_version_get());
            return Status::OK;
        }
};

class L2ServiceImpl final : public L2::Service {
    public:
        explicit L2ServiceImpl() {
        }
        Status Init(ServerContext* context, const InitRequest* request, InitResponse* response) {
            return Status(grpc::UNAVAILABLE, "");
        }
};

int main(int argc, char** argv) {
    std::string server_address("0.0.0.0:50051");
    DriverServiceImpl driverservice;
    L2ServiceImpl l2service;
    PortServiceImpl portservice;

    ServerBuilder builder;
    builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
    builder.RegisterService(&driverservice);
    builder.RegisterService(&portservice);
    builder.RegisterService(&l2service);
    std::unique_ptr<Server> server(builder.BuildAndStart());
    std::cout << "Server listening on " << server_address << std::endl;
    server->Wait();

    return 0;
}
