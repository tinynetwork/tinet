#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <linux/ip.h>
#include <arpa/inet.h>
#include <unistd.h>

int
create_listen_sock(struct sockaddr_in *addr)
{
  int error, sock;

  sock = socket(AF_INET, SOCK_STREAM, 0);
  if (sock == -1) {
    perror("socket");
    return -1;
  }

  error = bind(sock, (struct sockaddr *)addr, sizeof(*addr));
  if (error == -1) {
    perror("bind");
    return -1;
  }

  error = setsockopt(sock, SOL_IP, IP_TRANSPARENT, &(int){1}, sizeof(int));
  if (error == -1) {
    perror("setsockopt");
    return -1;
  }

  error = listen(sock, 100);
  if (error == -1) {
    perror("listen");
    return -1;
  }

  printf("listening...\n");

  return sock;
}

#define BUF_LEN 0xffff

void
serve(int lsock)
{
  int error;

  while (true) {
    int local_sock;
    struct sockaddr_in local_addr;
    socklen_t local_len = sizeof(local_addr);

    local_sock = accept(lsock, (struct sockaddr *)&local_addr, &local_len);
    if (local_sock == -1) {
      perror("accept");
      return;
    }

    printf("Local address %s:%u\n", inet_ntoa(local_addr.sin_addr), ntohs(local_addr.sin_port));

    int remote_sock;
    struct sockaddr_in remote_addr;
    socklen_t remote_len = sizeof(remote_addr);

    error = getsockname(local_sock, (struct sockaddr *)&remote_addr, &remote_len);
    if (error == -1) {
      perror("getpeername");
      return;
    }

    printf("Remote address %s:%u\n", inet_ntoa(remote_addr.sin_addr), ntohs(remote_addr.sin_port));

    remote_sock = socket(AF_INET, SOCK_STREAM, 0);
    if (remote_sock == -1) {
      perror("socket");
      return;
    }

    error = setsockopt(remote_sock, SOL_SOCKET, SO_REUSEADDR, &(int){1}, sizeof(int));
    if (error == -1) {
      perror("setsockopt");
      return;
    }

    error = setsockopt(remote_sock, SOL_IP, IP_TRANSPARENT, &(int){1}, sizeof(int));
    if (error == -1) {
      perror("setsockopt");
      return;
    }

    error = bind(remote_sock, (struct sockaddr *)&local_addr, local_len);
    if (error == -1) {
      perror("bind");
      return;
    }

    error = connect(remote_sock, (struct sockaddr *)&remote_addr, remote_len);
    if (error == -1) {
      perror("connect");
      return;
    }

    printf("Connected to %s:%u\n", inet_ntoa(remote_addr.sin_addr), ntohs(remote_addr.sin_port));

    uint8_t buf[BUF_LEN];
    ssize_t rlen, wlen, totlen = 0, ofs = 0;

    memset(buf, 0, BUF_LEN);

    rlen = read(local_sock, buf, BUF_LEN);
    if (rlen == -1) {
      perror("read");
      return;
    }

    printf("Forwarding request\n\n%s\n", buf);

    wlen = write(remote_sock, buf, rlen);
    if (wlen == -1) {
      perror("write");
      return;
    }

    memset(buf, 0, BUF_LEN);

    while (totlen < BUF_LEN) {
      rlen = read(remote_sock, buf + totlen, BUF_LEN - totlen);
      if (rlen == -1) {
        perror("read");
        return;
      } else if (rlen == 0) {
        break;
      }

      totlen += rlen;
    }

    printf("Forwarding reply\n\n%s\n", buf);

    ofs = 0;
    while (ofs < totlen) {
      wlen = write(local_sock, buf + ofs, totlen - ofs);
      if (wlen == -1) {
        perror("write");
        return;
      }

      ofs += wlen;
    }

    close(remote_sock);
    close(local_sock);
  }
}

int
main(void)
{
  int lsock;

  lsock = create_listen_sock(&(struct sockaddr_in){
    .sin_family = AF_INET,
    .sin_addr.s_addr = inet_addr("0.0.0.0"),
    .sin_port = htons(80),
  });
  if (lsock == -1) {
    fprintf(stderr, "sock_create failed\n");
    return EXIT_FAILURE;
  }

  serve(lsock);

  return EXIT_SUCCESS;
}
