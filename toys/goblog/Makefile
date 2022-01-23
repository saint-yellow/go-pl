# REMOTE=127.0.0.1
APPNAME=goblog

.PHONY: deploy
deploy:
	@echo "\n--- 开始构建可执行文件 ---"
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -o tmp/$(APPNAME)_tmp

	@echo "\n--- 上传可执行文件 ---"
	sudo cp tmp/$(APPNAME)_tmp /www/goblog.com

	@echo "\n--- 停止服务 ---"
	sudo supervisorctl stop $(APPNAME)

	@echo "\n--- 替换新文件 ---"
	cd /www/goblog.com/ && sudo rm $(APPNAME) && sudo mv $(APPNAME)_tmp $(APPNAME) && sudo chown root:root $(APPNAME)

	@echo "\n--- 开始服务 ---"
	sudo supervisorctl start $(APPNAME)

	@echo "\n--- 部署完毕 ---\n"
