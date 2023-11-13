# MutualExclusion
To run the program first ensure that the confFile.csv has exactly as many entries(rows) as the amount of peers that is going to be in the system.
Each row in the file needs an ip-address and a port seperated by a comma.

Next run the peer.go file and provide the row for the peer with -row and optionally a name for the peer with -name. The rows start at 0.

When the peers are running, type 'mutual' to send a request to the other peers for permission to access the critical section.
Type 'exit' to terminate
