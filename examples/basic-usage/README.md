# golcr client basic usage example

This describes a simple usage of the client to connect to the Leader and Clerk Resources on lds.org and print out a list of all the orgs to which the user has access.

To run, first modify the following lines in main.go:

```
const (
	user = "PUT_USERNAME_HERE"
	pass = "PUT_PASSWORD_HERE"
)
```

Replace `PUT_USERNAME_HERE` with the actual username and `PUT_PASSWORD_HERE` with the password for the user.

After that, run main.go.  The output will resemble the example output below (org names and numbers will vary).

```
Bishopric (123456)
Elders Quorum (123457)
Relief Society (123458)
Young Men (123459)
Young Women (123460)
Sunday School (123461)
Primary (123462)
Ward Missionaries (123463)
Full-Time Missionaries (123464)
Temple and Family History (123465)
Other Callings (123466)
```