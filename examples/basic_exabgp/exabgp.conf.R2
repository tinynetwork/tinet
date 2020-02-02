neighbor 10.0.0.1 {
        router-id 2.2.2.2;
        local-address 10.0.0.2;
        local-as 2;
        peer-as 1;
        #graceful-restart;

        static {
                route 8.8.8.8/32 next-hop 10.0.0.2;
        }
}
