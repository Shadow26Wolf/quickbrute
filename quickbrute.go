package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/leaanthony/clir"
	"github.com/liamg/tml"
)

// Output data from command
type Output struct {
	Host string
	Port int
}

var ftp []string
var ssh []string
var telnet []string
var smtp []string
var pop3 []string
var smb []string
var snmp []string
var ldap []string
var smb2 []string
var rexec []string
var rlogin []string
var rsh []string
var imap []string
var mssql []string
var oracle []string
var mysql []string
var rdp []string
var postgres []string
var vnc []string
var vnc2 []string
var irc []string

var verbose = false

func printLines(filePath string, values interface{}) error {
	if values == nil {
		return nil
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	rv := reflect.ValueOf(values)
	if rv.Kind() != reflect.Slice {
		return errors.New("Not a slice")
	}
	for i := 0; i < rv.Len(); i++ {
		fmt.Fprintln(f, rv.Index(i).Interface())
	}
	return nil
}

func main() {

	cli := clir.NewCli("Quick-Brute", "Automatic Service Bruteforce Tool", "v0.0.2")

	var inputDomains string
	var ports = "21,22,23,25,53,110,139,162,445,512,513,514,993,1433,1521,3306,3389,5432,5900,5901,6667"
	var threads = "10"
	var rate = "1000"
	cli.StringFlag("d", "Path to input domains to use", &inputDomains)
	cli.StringFlag("p", "Ports to scan", &ports)
	cli.StringFlag("t", "Number of concurrent goroutines for resolving", &threads)
	cli.StringFlag("rate", "Rate of scan probe requests", &rate)
	cli.BoolFlag("v", "Enable verbose output", &verbose)

	// Define action for the command
	cli.Action(func() error {
		if _, err := os.Stat(inputDomains); os.IsNotExist(err) {
			// path/to/whatever does not exist
			return err
		}

		tml.Println("[<blue>*</blue>] Starting Naabu...")

		commandArgs := []string{"-hL", inputDomains, "-ports", ports, "-json", "-Pn", "-t", threads, "-rate", rate}
		runCommand("naabu", commandArgs)

		printLines("/tmp/ftp.txt", ftp)
		printLines("/tmp/ssh.txt", ssh)
		printLines("/tmp/telnet.txt", telnet)
		printLines("/tmp/smtp.txt", smtp)
		printLines("/tmp/pop3.txt", pop3)
		printLines("/tmp/smb.txt", smb)
		printLines("/tmp/snmp.txt", snmp)
		printLines("/tmp/smb2.txt", smb2)
		printLines("/tmp/rexec.txt", rexec)
		printLines("/tmp/rlogin.txt", rlogin)
		printLines("/tmp/rsh.txt", rsh)
		printLines("/tmp/imap.txt", imap)
		printLines("/tmp/mssql.txt", mssql)
		printLines("/tmp/oracle.txt", oracle)
		printLines("/tmp/mysql.txt", mysql)
		printLines("/tmp/rdp.txt", rdp)
		printLines("/tmp/postgres.txt", postgres)
		printLines("/tmp/vnc.txt", vnc)
		printLines("/tmp/vnc2.txt", vnc2)
		printLines("/tmp/irc.txt", irc)

		// add commands for all services
		// add a check that checks if there are over ~200 results for a service
		// make config file that contains wordlist paths, download wordlist from github?
		// add the ability to run on a schedule
		// make it look pretty and try to make the above spam neater and faster :)

		if len(ftp) != 0 {
			tml.Printf("[<blue>*</blue>] Starting Medusa To Bruteforce %d FTP Services\n", len(ftp))
			commandArgs = []string{"-H", "/tmp/ftp.txt", "-M", "ftp", "-U", "Wordlist/ftp_u.txt", "-P", "Wordlist/ftp_p.txt", "-e", "ns", "-t", "20", "-T", "10"}
			runCommand("medusa", commandArgs)
		}

		if len(ssh) != 0 {
			tml.Printf("[<blue>*</blue>] Starting Medusa To Bruteforce %d SSH Services\n", len(ssh))
			commandArgs = []string{"-H", "/tmp/ssh.txt", "-M", "ssh", "-U", "Wordlist/ssh_u.txt", "-P", "Wordlist/ssh_p.txt", "-e", "ns", "-t", "20", "-T", "10"}
			runCommand("medusa", commandArgs)
		}

		if len(telnet) != 0 {
			tml.Printf("[<blue>*</blue>] Starting Medusa To Bruteforce %d Telnet Services\n", len(telnet))
			commandArgs = []string{"-H", "/tmp/telnet.txt", "-M", "telnet", "-U", "Wordlist/telnet_u.txt", "-P", "Wordlist/telnet_p.txt", "-e", "ns", "-t", "20", "-T", "10"}
			runCommand("medusa", commandArgs)
		}

		if len(smtp) != 0 {
			tml.Printf("[<blue>*</blue>] Starting Medusa To Bruteforce %d SMTP Services\n", len(smtp))
			commandArgs = []string{"-H", "/tmp/smtp.txt", "-M", "smtp", "-U", "Wordlist/smtp_u.txt", "-P", "Wordlist/smtp_p.txt", "-e", "ns", "-t", "20", "-T", "10"}
			runCommand("medusa", commandArgs)
		}

		if len(pop3) != 0 {
			tml.Printf("[<blue>*</blue>] Starting Medusa To Bruteforce %d POP3 Services\n", len(pop3))
			commandArgs = []string{"-H", "/tmp/pop3.txt", "-M", "pop3", "-U", "Wordlist/pop_u.txt", "-P", "Wordlist/pop_p.txt", "-e", "ns", "-t", "20", "-T", "10"}
			runCommand("medusa", commandArgs)
		}

		if len(smb) != 0 {
			tml.Printf("[<blue>*</blue>] Starting Medusa To Bruteforce %d SMB Services\n", len(smb))
			commandArgs = []string{"-H", "/tmp/smb.txt", "-M", "smbnt", "-U", "Wordlist/user.txt", "-P", "Wordlist/pass.txt", "-e", "ns", "-t", "20", "-T", "10"}
			runCommand("medusa", commandArgs)
		}

		if len(snmp) != 0 {
			tml.Printf("[<blue>*</blue>] Starting Medusa To Bruteforce %d SNMP Services\n", len(snmp))
			commandArgs = []string{"-H", "/tmp/snmp.txt", "-M", "snmp", "-U", "Wordlist/snmp.txt", "-P", "Wordlist/snmp.txt", "-e", "ns", "-t", "20", "-T", "10"}
			runCommand("medusa", commandArgs)
		}

		if len(smb2) != 0 {
			tml.Printf("[<blue>*</blue>] Starting Medusa To Bruteforce %d SMB2 Services\n", len(smb2))
			commandArgs = []string{"-H", "/tmp/smb2.txt", "-M", "smbnt", "-U", "Wordlist/user.txt", "-P", "Wordlist/pass.txt", "-e", "ns", "-t", "20", "-T", "10"}
			runCommand("medusa", commandArgs)
		}

		if len(rexec) != 0 {
			tml.Printf("[<blue>*</blue>] Starting Medusa To Bruteforce %d Rexec Services\n", len(rexec))
			commandArgs = []string{"-H", "/tmp/rexec.txt", "-M", "rexec", "-U", "Wordlist/user.txt", "-P", "Wordlist/pass.txt", "-e", "ns", "-t", "20", "-T", "10"}
			runCommand("medusa", commandArgs)
		}

		return nil
	})
	// Run the application
	err := cli.Run()
	if err != nil {
		// We had an error
		log.Fatal(err)
	}
}

func runCommand(command string, commandArgs []string) {
	cmd := exec.Command(command, commandArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			if verbose == true {
				tml.Printf("[<blue>*</blue>] <darkgrey>%s</darkgrey>\n", scanner.Text())
			}
			if command == "medusa" {
				if strings.Contains(scanner.Text(), "[SUCCESS]") {
					tml.Printf("[<green>+</green>] %s\n", scanner.Text())
				}
			}

			if command == "naabu" {
				outputJSON := scanner.Text()
				var output Output
				json.Unmarshal([]byte(outputJSON), &output)
				if output.Port == 21 {
					ftp = append(ftp, output.Host)
				}
				if output.Port == 22 {
					ssh = append(ssh, output.Host)
				}
				if output.Port == 23 {
					telnet = append(telnet, output.Host)
				}
				if output.Port == 25 {
					smtp = append(smtp, output.Host)
				}
				if output.Port == 110 {
					pop3 = append(pop3, output.Host)
				}
				if output.Port == 139 {
					smb = append(smb, output.Host)
				}
				if output.Port == 162 {
					snmp = append(snmp, output.Host)
				}
				if output.Port == 445 {
					smb2 = append(smb2, output.Host)
				}
				if output.Port == 512 {
					rexec = append(rexec, output.Host)
				}
				if output.Port == 513 {
					rlogin = append(rlogin, output.Host)
				}
				if output.Port == 514 {
					rsh = append(rsh, output.Host)
				}
				if output.Port == 993 {
					imap = append(imap, output.Host)
				}
				if output.Port == 1433 {
					mssql = append(mssql, output.Host)
				}
				if output.Port == 1521 {
					oracle = append(oracle, output.Host)
				}
				if output.Port == 3306 {
					mysql = append(mysql, output.Host)
				}
				if output.Port == 3389 {
					rdp = append(rdp, output.Host)
				}
				if output.Port == 5432 {
					postgres = append(postgres, output.Host)
				}
				if output.Port == 5900 {
					vnc = append(vnc, output.Host)
				}
				if output.Port == 5901 {
					vnc2 = append(vnc2, output.Host)
				}
				if output.Port == 6667 {
					irc = append(irc, output.Host)
				}
			}
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}
