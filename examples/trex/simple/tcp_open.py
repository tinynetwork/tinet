from trex.astf.api import *
import argparse


class Prof1():
    def __init__(self):
        pass  # tunables

    def create_profile(self, cps):
        prog_c = ASTFProgram()
        prog_c.connect()
        prog_c.reset()
        prog_s = ASTFProgram()
        prog_s.wait_for_peer_close()

        # ip generator
        ip_gen_c = ASTFIPGenDist(ip_range=["20.0.0.0", "20.0.255.255"], distribution="seq")
        ip_gen_s = ASTFIPGenDist(ip_range=["30.0.0.0", "30.0.255.255"], distribution="seq")
        ip_gen = ASTFIPGen(glob=ASTFIPGenGlobal(ip_offset="1.0.0.0"),
                           dist_client=ip_gen_c,
                           dist_server=ip_gen_s)

        # template
        temp_c = ASTFTCPClientTemplate(program=prog_c,  ip_gen=ip_gen, cps=cps)
        temp_s = ASTFTCPServerTemplate(program=prog_s)
        template = ASTFTemplate(client_template=temp_c, server_template=temp_s)

        # profile
        profile = ASTFProfile(default_ip_gen=ip_gen,
                              templates=template)
        return profile

    def get_profile(self, tunables, **kwargs):
        parser = argparse.ArgumentParser(
            description='Argparser for {}'.format(os.path.basename(__file__)),
            formatter_class=argparse.ArgumentDefaultsHelpFormatter)
        parser.add_argument('--cps', type=int, default=1)
        args = parser.parse_args(tunables)
        return self.create_profile(args.cps)


def register():
    return Prof1()
