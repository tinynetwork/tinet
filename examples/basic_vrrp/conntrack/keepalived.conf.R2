global_defs {
}
vrrp_instance VI_1 {
    state MASTER
    interface net0
    virtual_router_id 52
    priority 100
    advert_int 1
    virtual_ipaddress {
        10.0.0.254/24
    }
}
