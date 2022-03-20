default: build

build:
	go install

test1:
	echo 'var hq_str_sz002456="欧菲光,20.600,20.480,20.430,21.150,20.380,20.430,20.440,168464614,3488444639.620,28200,20.430,113500,20.420,132900,20.410,934900,20.400,142100,20.390,4800,20.440,55483,20.450,9200,20.460,20000,20.470,45300,20.480,2020-08-04,11:30:00,00";' | chart
test2:
	echo '5.9 2 2.3 1.8 7 5.7 5.5 13 1.3 3.4 2.3 23 5.8 0.6 3.8 21 2.4 7.6 4.8 7.2 0.7 5.5 6.5 2.4 8.8 4.7 2.1 1.8 7.6 3.3 1.4 4.5' | chart -d \ 
curl:
	curl http://hq.sinajs.cn/list\=sz002456 -s | chart

first:
	go-bindata -o assets.go -prefix . -pkg main tpl/*

assets:
	chart -a

timeline:
	echo '[{"timestamp": 1647760239, "value": 0.05 }, {"timestamp": 1647760240, "value": 0.03 }, {"timestamp": 1647760241, "value": 0.06 }, {"timestamp": 1647760242, "value": 0.1 }, {"timestamp": 1647760243, "value": 0.05 }, {"timestamp": 1647760244, "value": 0.05 }]' | chart -t timeline

tidy:
	go mod tidy