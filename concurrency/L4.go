package main

func addEmailsToQueue(emails []string) chan string {
	ansCh := make(chan string, len(emails))

	for _,v :=range(emails){
		ansCh<-v;
	}

	return ansCh
}
