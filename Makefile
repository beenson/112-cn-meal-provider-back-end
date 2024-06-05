.PHONY: proto api-gateway billing notification ordering rating user_mgmt

proto: common.pb.go user_mgmt.pb.go ordering.pb.go rating.pb.go notification.pb.go billing.pb.go

%.pb.go: proto/%.proto
	protoc --experimental_allow_proto3_optional --go_out=module=gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end:. \
		--go-grpc_out=module=gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end:. \
		-I proto $< 

service: api-gateway billing notification ordering rating user_mgmt

api-gateway: 
	cd src/api-gateway && go build -o ../../build/api-gateway
billing: 
	cd src/micro-svcs/billing && go build -o ../../../build/billing
notification: 
	cd src/micro-svcs/notification && go build -o ../../../build/notification
ordering: 
	cd src/micro-svcs/ordering && go build -o ../../../build/ordering
rating: 
	cd src/micro-svcs/rating && go build -o ../../../build/rating
user_mgmt: 
	cd src/micro-svcs/user_mgmt && go build -o ../../../build/user_mgmt

clean:
	rm -rf build/

docker-compose:
	cd docker/ && docker-compose up -d

docker-compose-rebuild:
	cd docker/ && docker-compose up -d --build --no-cache

docker-compose-down:
	cd docker/ && docker-compose down