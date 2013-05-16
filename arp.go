package main

/*
#include <unistd.h>
#include <sys/socket.h>
#include <netpacket/packet.h>
#include <net/ethernet.h>
#include <net/if_arp.h>
#include <sys/types.h>
#include <linux/if_ether.h>
#include <sys/ioctl.h>
#include <net/if.h>
#include <arpa/inet.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>
#include <signal.h>

#define EN_US_LANG

int recv_arp_fd;
#define ARP_RESPONSE 0x02
#define OFFSET_SOURCEIP 0x1c
#define OFFSET_ARPOP 0x14
#ifndef sys_err
#define sys_err(err_msg)                                        \
        do {                                                    \
                fprintf(stderr, "%s: %s,at line %d\n", err_msg, \
                        strerror(errno), __LINE__);             \
                return -1;                                      \
        } while (0)
#endif  // sys_err

#ifndef MAC_LEN
#define MAC_LEN 0x06
#endif  // MAC_LEN

typedef unsigned char u8_t;
typedef unsigned short u16_t;
typedef unsigned int u32_t;
typedef unsigned long u64_t;


typedef struct ehter_head {
        u8_t h_dest[MAC_LEN];
        u8_t h_source[MAC_LEN];
        u16_t h_proto;
} arp_frame_head_t;

struct ethernet_arp_head {
        u16_t hd_type;
        u16_t pr_type;
        u8_t hd_addr_len;
        u8_t pr_addr_len;
        u16_t op;
#if 0
        u8_t src_mac[MAC_LEN];
        u8_t src_ip[4];
        u8_t dest_mac[MAC_LEN];
        u8_t dest_ip[4];
#endif
};

typedef struct ethernet_arp {
        struct ethernet_arp_head eah;
        u8_t src_mac[MAC_LEN];
        u8_t src_ip[4];
        u8_t dest_mac[MAC_LEN];
        u8_t dest_ip[4];
} arp_frame_data_t;

#define hd_type eah.hd_type
#define pr_type eah.pr_type
#define hd_addr_len eah.hd_addr_len
#define pr_addr_len eah.pr_addr_len
#define op eah.op

typedef struct arp {
        arp_frame_head_t eth_header;
        arp_frame_data_t arp_header;
        u8_t padding[18];
} arp_t;


static inline int get_ip_addr(const char *interface, u32_t *ip)
{
        struct ifreq ipreq;
        int ip_fd;
        struct sockaddr_in *sin;

        if (interface == NULL)
                return -1;
        if ((ip_fd = socket(AF_INET, SOCK_DGRAM, 0)) < 0)
#ifdef EN_US_LANG
                sys_err("When open socket file description failed");
#else
                sys_err("打开套接字失败");
#endif

        bzero(&ipreq, sizeof(ipreq));
        strcpy(ipreq.ifr_name, interface);

        if (ioctl(ip_fd, SIOCGIFADDR, &ipreq) < 0)
#ifdef EN_US_LANG
                sys_err("When get ip address failed");
#else
                sys_err("获取IP地址失败");
#endif
        sin = (struct sockaddr_in *)&ipreq.ifr_addr;

        memcpy(ip, &sin->sin_addr.s_addr, 4);
        close(ip_fd);

        return 0;
}

static inline int get_mac_addr(const char *interface, u8_t mac[])
{
        struct ifreq req;
        int mac_fd;

        if (interface == NULL)
                return -1;
        if ((mac_fd = socket(AF_INET, SOCK_DGRAM, 0)) < 0)
#ifdef EN_US_LANG
                sys_err("When open socket file description failed");
#else
                sys_err("打开套接字失败");
#endif

        bzero(&req, sizeof(req));

        strcpy(req.ifr_name, interface);
        if (ioctl(mac_fd, SIOCGIFHWADDR, &req) < 0)
#ifdef EN_US_LANG
                sys_err("When get MAC address failed");
#else
                sys_err("获取MAC失败");
#endif

        memcpy(mac, req.ifr_hwaddr.sa_data, MAC_LEN);
        close(mac_fd);

        return 0;
}

static int prip(const void *ip)
{
        char ip_buf[24];

        if (inet_ntop(AF_INET, ip, ip_buf, sizeof(ip_buf)) == NULL)
#ifdef EN_US_LANG
                sys_err("When convert ip to string failed");
#else
                sys_err("转换IP地址失败");
#endif

        puts(ip_buf);

        return 0;
}


static void prmac(const u8_t *mac)
{
        printf("%02x:%02x:%02x:%02x:%02x:%02x\n",
               *mac,*(mac+1),*(mac+2),*(mac+3),*(mac+4),*(mac+5));
}

static inline int print_ip_mac(const void *ip, const u8_t *mac)
{
        printf("IP: ");
        if (prip(ip) < 0)
        	return -1;
        puts("\t/\\");
        puts("\t||");
        puts("\t\\/");
        printf("MAC: ");
        prmac(mac);

        return 0;
}

void print_arp_frame(const arp_t *p)
{
#ifdef EN_US_LANG
        puts("Structure of arp is:");
        puts("\tHeader:");
        printf("\t\tDestination MAC address: ");
        prmac(p->eth_header.h_dest);
        printf("\t\tSource MAC address: ");
        prmac(p->eth_header.h_source);
        printf("\t\tFrame type: %04x\n", ntohs(p->eth_header.h_proto));
        puts("\tARP Request or Reply:");
        printf("\t\tHardware type: %04x\n", ntohs(p->arp_header.hd_type));
        printf("\t\tProtocol type: %04x\n", ntohs(p->arp_header.pr_type));
        printf("\t\tHardware address length: %02x\n", p->arp_header.hd_addr_len);
        printf("\t\tProtocol address length: %02x\n", p->arp_header.pr_addr_len);
        printf("\t\tRequest or reply: %04x\n", ntohs(p->arp_header.op));
        printf("\t\tSource MAC address: ");
        prmac(p->arp_header.src_mac);
        printf("\t\tSource ip address: ");
        prip(p->arp_header.src_ip);
        printf("\t\tDestination MAC address: ");
        prmac(p->arp_header.dest_mac);
        printf("\t\tDestination ip address: ");
        prip(p->arp_header.dest_ip);
        printf("\t\tFill %zu bytes\n", sizeof(p->padding));
        puts("End");
#else
        puts("ARP帧内容:");
        puts("\tARP帧头:");
        printf("\t\t目的MAC地址: ");
        prmac(p->eth_header.h_dest);
        printf("\t\t源MAC地址: ");
        prmac(p->eth_header.h_source);
        printf("\t\t帧类型: %04x\n", ntohs(p->eth_header.h_proto));
        puts("\tARP帧内容:");
        printf("\t\t硬件类型: %04x\n", ntohs(p->arp_header.hd_type));
        printf("\t\t协议类型: %04x\n", ntohs(p->arp_header.pr_type));
        printf("\t\t硬件地址长度: %02x\n", p->arp_header.hd_addr_len);
        printf("\t\t协议地址长度: %02x\n", p->arp_header.pr_addr_len);
        printf("\t\t操作类型: %04x\n", ntohs(p->arp_header.op));
        printf("\t\t源MAC地址: ");
        prmac(p->arp_header.src_mac);
        printf("\t\t源IP地址: ");
        prip(p->arp_header.src_ip);
        printf("\t\t目的MAC地址: ");
        prmac(p->arp_header.dest_mac);
        printf("\t\t目的IP地址: ");
        prip(p->arp_header.dest_ip);
        printf("\t\t填充字节数%zu(字节)\n", sizeof(p->padding));
        puts("结束");
#endif
}

void timeout(int sig)
{
#ifdef EN_US_LANG
        fprintf(stderr, "Time OUT !\n");
#else
        fprintf(stderr, "超时 !\n");
#endif
        exit(1);
}

int arp(char *ip, char *interface)
{
        int send_arp_fd;
        struct sockaddr_ll other;
        struct sockaddr_ll from;
        arp_t send_arp;
        arp_t *recv_arp = NULL;
        socklen_t len;
        // u8_t *ptr = NULL;
        int n;
        u32_t local_ip;
        char  recv_buf[120];

        // ---------------------------------Create send socket-------------------------------------
        send_arp_fd = socket(AF_PACKET, SOCK_RAW, htons(ETH_P_ARP));
        if (send_arp_fd < 0)
#ifdef EN_US_LANG
                sys_err("When create socket fd failed");
#else
                sys_err("打开套接字失败");
#endif

        if (setuid(getuid()) < 0) {
#ifdef EN_US_LANG
                sys_err("When set uid failed");
#else
                sys_err("设置uid失败");
#endif
        }
        // ----------------------------------Fill Datagram-------------------------------------
        get_mac_addr(interface, send_arp.eth_header.h_source);
        get_mac_addr(interface, send_arp.arp_header.src_mac);

        memset(send_arp.eth_header.h_dest, 0xFF, sizeof(send_arp.eth_header.h_dest));
        memset(send_arp.padding, 0, sizeof(send_arp.padding)); // fill 0 to 60 bytes

        send_arp.eth_header.h_proto = htons(ETH_P_ARP);
#if 0
        ptr = (u8_t *)&send_arp.eth_header.h_proto;
        printf("帧类型###########%02x###%02x################\n", ptr[0], ptr[1]);
#endif
        send_arp.arp_header.hd_type = htons(ARPHRD_ETHER);
#if 0
        ptr = (u8_t *)&send_arp.arp_header.hd_type;
        printf("硬件类型###########%02x###%02x################\n", ptr[0], ptr[1]);
#endif
        send_arp.arp_header.pr_type = htons(ETH_P_IP);
#if 0
        ptr = (u8_t *)&send_arp.arp_header.pr_type;
        printf("协议类型###########%02x###%02x################\n", ptr[0], ptr[1]);
#endif
        send_arp.arp_header.op = htons(ARPOP_REQUEST);

        send_arp.arp_header.hd_addr_len = 6;
        send_arp.arp_header.pr_addr_len = 4;
        get_ip_addr(interface, &local_ip);
        memcpy(send_arp.arp_header.src_ip, &local_ip, 4);


        if (inet_pton(AF_INET, ip, &send_arp.arp_header.dest_ip) < 0)
#ifdef EN_US_LANG
                sys_err("When convert ip to network format failed");
#else
                sys_err("转换IP地址失败");
#endif

        memset(send_arp.arp_header.dest_mac, 0xFF,
               sizeof(send_arp.arp_header.dest_mac));

        // ----------------------------------Send arp frame-------------------------------------
        bzero(&other, sizeof(other));
        other.sll_family = AF_PACKET;
        other.sll_protocol = htons(ETH_P_ARP);

        other.sll_ifindex = if_nametoindex(interface);
        // other.sll_pkttype = PACKET_BROADCAST;
        // other.sll_hatype = ARPHRD_ETHER;
        // other.sll_halen = ETH_ALEN;
        // memset(other.sll_addr, 0xFF, sizeof(other.sll_addr));
#ifdef EN_US_LANG
        printf("Before Send !!!\n");
#else
        printf("发送之前 !!!\n");
#endif
        print_arp_frame(&send_arp);
        if (sendto(send_arp_fd, &send_arp, sizeof(send_arp),
                   0, (struct sockaddr *)&other, sizeof(other)) < 0)
#ifdef EN_US_LANG
                sys_err("When send frame data failed");
#else
                sys_err("发送数据失败");
#endif


        // ---------------------------------Create receive socket-------------------------------------
        recv_arp_fd = socket(AF_PACKET, SOCK_RAW, htons(ETH_P_ARP));
        if (recv_arp_fd < 0)
#ifdef EN_US_LANG
                sys_err("When create socket fd failed");
#else
                sys_err("打开套接字失败");
#endif

        // ----------------------------------Receive arp frame-------------------------------------
        bzero(recv_buf, sizeof(recv_buf));
        bzero(&from, sizeof(from));
        len = sizeof(from);
        signal(SIGALRM, timeout);
        alarm(60);
        while (1) {
#ifdef EN_US_LANG
                printf("After Send !!!\n");
#else
                printf("发送之后 !!!\n");
#endif
                if ((n = recvfrom(recv_arp_fd, recv_buf,
                                  sizeof(recv_buf), 0, (struct sockaddr *)&from, &len)) < 0)
#ifdef EN_US_LANG
                        sys_err("When receive frame data failed");
#else
                        sys_err("接收数据失败");
#endif
                recv_arp = (arp_t *)recv_buf;
                print_arp_frame(recv_arp);
#ifdef EN_US_LANG
                printf("%d bytyes received\n", n);
#else
                printf("接收[%d]字节\n", n);
#endif
                if ((ntohs(*(uint16_t *)(recv_buf + OFFSET_ARPOP)) == ARP_RESPONSE) &&
                    (memcmp(&send_arp.arp_header.dest_ip, recv_buf + OFFSET_SOURCEIP, 4) == 0))
                        break;
        }

        close(recv_arp_fd);
        close(send_arp_fd);
        if (print_ip_mac(recv_arp->arp_header.src_ip, recv_arp->arp_header.src_mac) < 0)
        	return -1;
        return 0;
}
*/
import "C"

import (
        "flag"
        "unsafe"
)

func main() {
        var go_ip string
        var go_inter string
        flag.StringVar(&go_ip, "ip", "", "对方IP地址")
        flag.StringVar(&go_inter, "interface", "eth0", "物理接口名称")
        help := flag.Bool("help", false, "显示帮助")
        help = flag.Bool("h", false, "显示帮助")
        flag.Parse()
        if *help || go_ip == "" {
                flag.PrintDefaults()
                return
        }
        ip := C.CString(go_ip)
        inter := C.CString(go_inter)
        defer C.free(unsafe.Pointer(ip))
        defer C.free(unsafe.Pointer(inter))
        if C.arp(ip, inter) != 0 {

        }
}
