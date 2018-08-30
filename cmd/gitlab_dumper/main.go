package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	json_definition "github.com/PapaYofen/gitlab_dumper/json"
)

var (
	url      = flag.String("url", "http://gitlab.example.com", "gitlab url")
	token    = flag.String("token", "aG64MKstv8sW5VZRXwL9", "private token")
	groupSrc = flag.String("src-group", "src", "group name")
	groupDst = flag.String("dst-group", "dst", "group name")
	user     = flag.String("user", "yofan", "user")
	passwd   = flag.String("passwd", "yofan", "password")
)

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body, nil
}

func httpPost(url string) ([]byte, error) {
	resp, err := http.Post(url, "", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body, nil
}

func listGroups(url string, targets []string) ([]string, error) {
	if len(targets) != 2 {
		return nil, errors.New("target slice size must be 2")
	}
	body, _ := httpGet(url)
	//	parse
	var groups json_definition.Groups
	if err := json.Unmarshal(body, &groups); err != nil {
		fmt.Println("maybe recv Unauthorized message, pls check resp body below")
		fmt.Println(string(body))
		panic(err)
		// return nil, err
	}
	ids := make([]string, len(targets))
	for _, g := range groups {
		for idx, t := range targets {
			if g.Name == t && g.FullName == t {
				ids[idx] = strconv.Itoa(g.ID)
				fmt.Printf("===> id #%d=%s, group name=%s\n", idx, ids[idx], t)
			}
		}
	}
	if len(ids[0]) == 0 || len(ids[1]) == 0 {
		panic("NOT found targets group id")
		// return nil, errors.New("NOT found targets group id")
	}
	return ids, nil
}

func listProjects(url string) ([]string, []string, error) {
	body, _ := httpGet(url)
	//	parse
	var projects json_definition.Projects
	if err := json.Unmarshal(body, &projects); err != nil {
		panic(err)
		// return nil, nil, err
	}
	names := make([]string, len(projects))
	urls := make([]string, len(projects))
	for idx, p := range projects {
		names[idx] = p.Name
		urls[idx] = p.WebURL
		fmt.Printf("===> name=%s,url=%s\n", names[idx], urls[idx])
	}
	return names, urls, nil
}

func main() {
	flag.Parse()

	//	get id of src and dst groups
	groups := []string{*groupSrc, *groupDst}
	requestURL := *url + "/api/v4/groups?private_token=" + *token
	fmt.Printf("list group url=%s\n", requestURL)
	ids, err := listGroups(requestURL, groups)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n")

	//	get projects under src group
	requestURL = *url + "/api/v4/groups/" + ids[0] + "/projects?private_token=" + *token
	fmt.Printf("list project url=%s\n", requestURL)
	names, urls, err := listProjects(requestURL)

	fmt.Printf("\n")

	//	create project under dst group and
	//	import external project from src group to project created
	for idx, _ := range names {
		importURL := "http://" + *user + ":" + *passwd + "@" + urls[idx][len("http://"):]
		requestURL = *url + "/api/v4/projects?private_token=" + *token +
			"&namespace_id=" + ids[1] +
			"&name=" + names[idx] +
			"&import_url=" + importURL
		fmt.Println("post url=", requestURL)
		httpPost(requestURL)
	}
}
