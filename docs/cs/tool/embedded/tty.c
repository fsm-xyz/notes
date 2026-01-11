#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <termios.h>
#define DEV_NAME "/dev/ttyUSB0"
int main(int argc, char *argv[])
{
    int fd;
    int len, i, ret;

    fd = open(DEV_NAME, O_RDWR | O_NOCTTY);
    if (fd < 0)
    {
        perror(DEV_NAME);
        return -1;
    }
    struct termios uart_cfg_opt;
    speed_t speed = B38400;
    if (-1 == tcgetattr(fd, &uart_cfg_opt))
        return -1;
    tcflush(fd, TCIOFLUSH);
    cfsetospeed(&uart_cfg_opt, speed);
    cfsetispeed(&uart_cfg_opt, speed);
    if (-1 == tcsetattr(fd, TCSANOW, &uart_cfg_opt))
        return -1;
    uart_cfg_opt.c_cc[VTIME] = 1;
    uart_cfg_opt.c_cc[VMIN] = 0;
    /* Data length setting section */ uart_cfg_opt.c_cflag &= ~CSIZE;
    uart_cfg_opt.c_cflag |= CS8;
    uart_cfg_opt.c_iflag &= ~INPCK;
    uart_cfg_opt.c_cflag &= ~PARODD;
    uart_cfg_opt.c_cflag &= ~CSTOPB;
    /* Using raw data mode */
    uart_cfg_opt.c_lflag &= ~(ICANON | ECHO | ECHOE | ISIG);
    uart_cfg_opt.c_iflag &= ~(INLCR | IGNCR | ICRNL | IXON | IXOFF);
    uart_cfg_opt.c_oflag &= ~(INLCR | IGNCR | ICRNL);
    uart_cfg_opt.c_oflag &= ~(ONLCR | OCRNL);
    /* Apply new settings */
    if (-1 == tcsetattr(fd, TCSANOW, &uart_cfg_opt))
        return -1;
    tcflush(fd, TCIOFLUSH);

    unsigned char buf[] = {0x01, 0x06, 0x00, 0x01, 0x00, 0x01, 0x48, 0x0A};

    printf("buf:");
    for (i = 0; i < sizeof(buf); i++)
    {
        printf("0x%02x ", *(buf + i));
    };
    printf("\n");

    len = write(fd, buf, sizeof(buf)); /* 向串口写入字符串 */
    if (len < 0)
    {
        printf("write data error \n");
    }
    printf("write data ok, len: %d \n", len);

    len = read(fd, buf, sizeof(buf)); /* 在串口读入字符串 */
    printf("read data ok, len: %d \n", len);
    if (len < 0)
    {
        printf("read error \n");
        return -1;
    }

    printf("buf:");
    for (i = 0; i < sizeof(buf); i++)
    {
        printf("0x%02x ", *(buf + i));
    };
    printf("\n");
    return (0);
}
