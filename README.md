# 使用方法：

### 读取数据

  	yamlFile, err := ioutil.ReadFile("docker-compose.yml")
	if err != nil {
	    log.Printf("yamlFile.Get err #%v ", err)
	}

	conf := new(docker.Yml)
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
	    log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(conf.Version)
	fmt.Println(conf.Services["db"].Environment)
  
  
### 修改数据

	d, err := yaml.Marshal(conf)
	if err != nil {
	    log.Fatalf("error: %v", err)
	}
	ymlFile, err := os.OpenFile('test.yml', os.O_WRONLY, os.ModePerm)
	if err != nil {
	    fmt.Printf("permission denied![%v]\n", err)
	}

	_, err = ymlFile.WriteString(string(d))

	if err != nil {
	    fmt.Printf("permission denied![%v]\n", err)
	}
