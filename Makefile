default: build

build:
	go install
	
line:
	echo '5.9 2 2.3 1.8 7 5.7 5.5 13 1.3 3.4 2.3 23 5.8 0.6 3.8 21 2.4 7.6 4.8 7.2 0.7 5.5 6.5 2.4 8.8 4.7 2.1 1.8 7.6 3.3 1.4 4.5' | chart -t line -d \ 

timeline:
	echo '[{"timestamp": 1647760239, "value": 0.05 }, {"timestamp": 1647760240, "value": 0.03 }, {"timestamp": 1647760241, "value": 0.06 }, {"timestamp": 1647760242, "value": 0.1 }, {"timestamp": 1647760243, "value": 0.05 }, {"timestamp": 1647760244, "value": 0.05 }]' | chart -t timeline

tidy:
	go mod tidy