package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)
type Company struct {
	name string
	link string
	country string
	part string
	value string
}

var baseUrl string = "https://www.cbinsights.com/research-unicorn-companies"

func main(){
	list := getNames();
	writeFile(list);
}
func writeFile(list []Company){
	file,err := os.Create("unicorns.csv");
	checkErr(err);
	
	writer := csv.NewWriter(file);
	defer writer.Flush();
	
	header := []string{"Name","Country","Domain","Value","Link"};
	err = writer.Write(header);
	checkErr(err);
	
	for _,val := range list {
		comSlice := []string {val.name, val.country, val.part, val.value, val.link}
		err := writer.Write(comSlice);
		checkErr(err)
	}
	return;
}
func getNames()[]Company{
	var list = []Company{};
	var c = make(chan Company);
	res, err := http.Get(baseUrl);
	checkErr(err);
	checkStatus(res);

	defer res.Body.Close();

	doc, err := goquery.NewDocumentFromReader(res.Body);
	checkErr(err);
	coms := doc.Find("tr")
	coms.Each(func(i int, s *goquery.Selection){
		go extractCom(s, c)
	});
	for i := 0; i < coms.Length(); i++ {
		list = append(list, <-c);

	}
	return list;
}
func extractCom(s *goquery.Selection , c chan<-Company){
	tmp := Company{}
	s.Find("td").Each(func(i int, s *goquery.Selection){
		if i ==0 {
			link,_ := s.Children().Attr("href");
			name :=s.Children().Text();
			tmp.link = link;
			tmp.name = strings.Replace(name,"&amp;","&",-1);
		} else if i==1{
			value :=s.Text();
			tmp.value = value;
		} else if i==3{
			country :=s.Text();
			tmp.country = country;
		} else if i==5{
			part := s.Text();
			tmp.part = strings.Replace(part,"&amp;","&",-1);;
		}
	});
	c<- tmp;
	return;
}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err);
	}
}
func checkStatus (res *http.Response){
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed");
	}
}