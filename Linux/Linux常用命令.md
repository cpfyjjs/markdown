# Linux常用命令



[TOC]

***

## 学前理论

- **linux主要特征** ：一切且文件（目录、硬盘等都是文件）；硬件都在/dev 目录，如硬盘、U盘为/dev/sd[a-d]； /dev/sr0（/dev/cdrom）是光驱的设备名（df命令查看），为设备文件，代表的是光驱本身，得把这个设备挂载到目录下（一般为/mnt）(文件系统的临时挂载点)，才能对设备上的文件进行读写等操作；
- **不懂的命令** ：man 命令（查用法、全称），只记得关键词，可用man -k 关键词；
- **Linux命令常用结构** ：`Command [-option] [argument]`
  Command：即是要运行的命令的本身，说白了就是一个软件（程序）；
  Option：是选项（可选），选项是控制命令运行状态和行为的（可多个选项一起，如df -hT）；
  Argument：是参数（可选），是命令要操作对象如文件、路径、数据、目录等；
  在指令的第一部分按[tab]键一下为[命令补全]，两下为所有命令选择，在非第一部分按[tab]键两下为[文件补全]；
- **linux命令区分大小写**；

## 开关机

- **sync** ：把内存中的数据写到磁盘中（关机、重启前都需先执行sync）
- **shutdown -r now**或**reboot** ：立刻重启
- **shutdown -h now** ：立刻关机
- **shutdown -h 20:00** ：预定时间关闭系统（晚上8点关机，如果现在超过8点，则明晚8点）
- **shutdown -h +10** ：预定时间关闭系统（10分钟后关机）
- **shutdown -c** ：取消按预定时间关闭系统

## 系统信息

- **who am i** ：查看当前使用的终端
- **who** 或 **w** ： 查看所有终端
- **uname -m** ：显示机器的处理器架构（如x86_64）
- **cat /proc/version** ：查看linux版本信息
- **uname -r** ：显示正在使用的内核版本
- **rpm -qa | grep kernel-devel** ：查看kernel-devel版本（安装软件时编译内核用，故需要保持内核版本一致性）
- **yum install -y "kernel-devel-uname-r == $(uname -r)"**：安装和Linux内核版本匹配的kernel-devel
- **date** ：显示系统日期 （date +%Y/%m/%d : 显示效果如2018/01/01）
- **date 070314592018.00** ：设置时间（格式为月日时分年.秒 ）
- **clock -w** ：将时间修改保存到 BIOS
- **cal 2018** ：显示2018年的日历表
- **clear** ：清空命令行
- **ifconfig** ：显示或设置网卡（查ip等）（类似windows中ipconfig）
- **ping -c 3 www.baidu.com** ：测试百度与本机的连接情况（ -c 3表示测试3次）
- **cat /proc/cpuinfo** ：显示CPU的信息
- **cat /proc/cpuinfo| grep "physical id"| sort| uniq| wc -l** ：查看物理CPU个数
- **cat /proc/cpuinfo| grep "cpu cores"| uniq** ：查看每个物理CPU的核数
- **cat /proc/cpuinfo| grep "processor"| wc -l** ：查看逻辑CPU个数即线程数

## 系统性能

- **top** ：动态实时显示cpu、内存、进程等使用情况（类似windows下的任务管理器）
- **top -d 2 -p 7427** ：-d为画面更新的秒数，默认5秒，-p为指定进程pid的信息
- **vmstat 2 10** ：每隔2秒采集一次服务器状态，采集10次（查看内存、io读写状态、cpu）
- **free -h** :查看系统内存及虚拟内存使用情况
- **df -h** :显示磁盘的空间使用情况
- **iostat** ：可查io读写、cpu使用情况
- **sar -u 3 5** :查看cpu使用情况（3秒一次，共5次）
- **sar -d 2 3** ：评估磁盘性能
- **ps aux|grep firefox** ：获取火狐的进程号（PID）（可查看进程占用cpu、内存百分比及进程触发指令的路径）
- **kill -9 进程号** ：强制杀死进程

## 文件和目录

>  cd:是Change Directory的缩写，用来切换工作目录，语法：cd [相对或绝对路径或特殊符号]

- **cd** ：进入该用户的主目录 ~（root用户为/root,其他用户为/home/用户名）

- **cd ..** ：返回上一级目录（注意要空格）

- **cd -** ：返回上次所在目录

- **cd /** ：返回根目录 （绝对路径）

- **cd ./目录1/目录2** ：进入当前目录下的子目录（相对路径）

- **pwd** ：显示工作路径（Print Working Directory 的缩写）

  > ls:是List的缩写，用于列出目录下的文件，语法：`ls [选项][目录或文件名]`

- **ls -a** :列出文件下所有的文件，包括以“.“开头的隐藏文件

- **ls -lh *.log** :列出文件的详细信息（.log结尾，*为通配符代表任意多个字符）

- **file 文件或目录** ：显示文件的类型（目录、text、zip、shell脚本等）

- **mkdir dir1** :创建目录(dir1)（mkdir为make directory的缩写）

- **mkdir -p ./dir1/dir2** :递归创建目录（-p：父目录不存在时，同时建立）

- **touch a.txt** :创建文件a.txt

  > rm:可以删除一个目录中的一个或多个文件或目录，也可以将某个目录及其下属的所有文件及其子目录均删除掉; 语法：rm (选项)(参数)（注：如果参数中含有目录，则必须加上-r选项）；

- **rm 文件** ：删除文件

- **rm -r 目录或文件** ：删除目录（及目录下所有文件）（非空也可以）

- **rm -rf 目录或文件** ：强制删除，如：rm -rf * 为删除当前目录下所有文件

  > mv：是move的缩写，可以用来剪切移动文件、目录或者将文件改名；
  > 语法：mv 源文件 目标文件（改名）或目录（移动）；

- **mv a b** :移动或者重命名一个文件或者目录（存在即移动目录或覆盖文件，不存在即改名）

- **mv /opt/git/g /opt/a** ：移动g到opt目录下并改名为a（a目录不存在，若存在则为移动g到a目录下）

- **mv -t ./test a.txt b.txt** ：移动多个文件到某目录下

  > cp:复制文件或目录；cp命令可以将单个或多个文件复制到一个已经**存在**的目录下；
  > 常用：cp -ai 文件或目录 目标目录;

- **cp -ai /opt/abc /opt/git/** ：复制abc目录（或文件）到git目录下（选项a表示文件的属性也复制、目录下所有文件都复制；i表示覆盖前询问）

  > ln：link的缩写，用于建立硬（软）链接，常用于软件安装时建软链接(类似快捷方式)到PATH;
  > 语法：ln [-s] 源文件 目标文件

- **ln -s /opt/a.txt /opt/git/** :对文件创建软链接（快捷方式不改名还是a.txt）

- **ln -s /opt/a.txt /opt/git/b** :（快捷方式改名为b）（下面的一样可以改名）

- **ln -s /opt/mulu /opt/git/** :对目录创建软链接

- **ln /opt/a.txt /opt/git/** :对文件创建硬链接

## 文件权限

- **chmod [-R] 777文件或目录** ：设置权限（chmod a+rwx a=chmod ugo +rwx a=chmod 777 a）

  > 注： r（read）对应4，w（write）对应2，x（execute）执行对应1；
  > -R：递归更改文件属组，就是在更改某个目录文件的属组时，如果加上-R的参数，那么该目录下的所有文件的属组都会更改）

- **chmod [{ugoa}{+-=}{rwx}][文件或目录]** ：如chmod u-w,g+x,o=r test.txt为user（拥有者）去掉写权限，group(所属组)加上执行权限，other(其他人)权限等于只读；

- **chown [-R] admin:root /opt/** ：变更文件及目录的拥有者和所属组（-R递归处理所有文件和文件夹，admin为拥有者，root为所属者）

## 文件查找

- **locate a.txt** ：在系统全局范围内查找文件名包含a.txt字样的文件（比find快）;

> locate:原理是updatedb会把文件系统中的信息存放到数据库databases中（但一般一天才执行一次，所以locate找不到新创建的文件，需要先手动执行updatedb，再执行locate）,locate从数据库中读数据;

>  find：在目录结构中搜索文件，并执行指定的操作
> 语法：find pathname -options [-print -exec ...]
> pathname ：为 find命令所查找的目录路径。例如用.来表示当前目录，用/来表示系统根目录（find查找范围为目标目录及其子目录所有文件及目录）；
> -exec： find命令对匹配的文件执行该参数所给出的shell命令。相应命令的形式为'command' { } ;，注意{ }和；之间的空格；
> -print： find命令将匹配的文件输出到标准输出；

- **find /home -mtime -2** ：在/home下查最近2*24小时内改动过的文件
- **find . -size +100M** ：在当前目录及子目录下查找大于100M的文件
- **find . -type f** ：f表示文件类型为普通文件（b/d/c/p/l/f 分别为块设备、目录、字符设备、管道、符号链接、普通文件）
- **find . -mtime +2 -exec rm {} ;** :查出更改时间在2*24小时以前的文件并删除它**
- **find . -name '\*.log' -exec grep -i hello {} \; -print** :在当前目录及子目录下查出文件名后缀为.log的文件并且该文件内容包含了hello字样并打印，-exec 命令 {} \表示对查出文件操作，-i表示不区分大小写；
- **find . -name '\*.log'|grep hello** :在当前目录及子目录下查出文件名后缀为.log的文件并且文件名包含了hello字样（grep用来处理字符串）；
- **grep -i 'HELLO' . -r -n** ：在当前目录及子目录下查找文件内容中包含hello的文件并显示文件路径（-i表示忽略大小写）
- **which java** ：在环境变量$PATH设置的目录里查找符合条件的文件，并显示路径（查询运行文件所在路径）
- **whereis java** :查看安装的软件的所有的文件路径（whereis 只能用于查找二进制文件、源代码文件和man手册页，一般文件的定位需使用locate命令）

## 查看文件的内容

- **cat [-n] 文件名** :显示文件内容，连行号一起显示
- **less 文件名** ：一页一页的显示文件内容（搜索翻页同man命令）
- **head [-n] 文件名** ：显示文件头n行内容，n指定显示多少行
- **tail [-nf] 文件名**:显示文件尾几行内容,n指定显示多少行,f用于实时追踪文件的所有更新，常用于查阅正在改变的日志文件（如tail -f -n 3 a.log 表示开始显示最后3行，并在文件更新时实时追加显示，没有-n默认10行）
- **sed -n '2,$p' ab** ：显示第二行到最后一行；
- **sed -n '/搜索的关键词/p' a.txt** ：显示包括关键词所在行
- **cat filename |grep abc -A10** ：查看filename中含有abc所在行后10行（A10）、前10行（B10）内容
- **less a.txt|grep git** ：显示关键词所在行，管道符”|”它只能处理由前面一个指令传出的正确输出信息，对错误信息信息没有直接处理能力。然后传递给下一个命令，作为标准的输入；
- **cat /etc/passwd |awk -F ':' '{print $1}'** ：显示第一列

## 文本处理

- **ls -l>file** ：输出重定向>（改变原来系统命令的默认执行方式）：ls -l命令结果输出到file文件中，若存在，则覆盖
- **cat file1 >>file** ：输出重定向之cat命令结果输出追加到file文件
- **ls fileno 2>file** ： 2>表示重定向标准错误输出（文件不存在，报错信息保存至file文件）；
- **cowsay <a.txt** :重定向标准输入’命令<文件’表示将文件做为命令的输入（为从文件读数据作为输入）
- **sed -i '4,$d' a.txt** ：删除第四行到最后一行（$表示最后一行）（sed可以增删改查文件内容）
- **sed -i '$a 增加的字符串' a.txt** ：在最后一行的下一行增加字符串
- **sed -i 's/old/new/g' a.txt** :替换字符串；格式为sed 's/要替换的字符串/新的字符串/g' 修改的文件
- **vim 文件**：编辑查看文件（同vi）

## 用户与权限

- **useradd 用户名** ：创建用户
- **userdel -r 用户名** :删除用户：（-r表示把用户的主目录一起删除）
- **usermod -g 组名 用户名** ：修改用户的组
- **passwd [ludf] 用户名** ：用户改自己密码，不需要输入用户名，选项-d:指定空口令,-l:禁用某用户，-u解禁某用户，-f：强迫用户下次登录时修改口令
- **groupadd 组名** ：创建用户组
- **groupdel 用户组** ：删除组
- **groupmod -n 新组名 旧组名** ：修改用户组名字
- **su - 用户名**：完整的切换到一个用户环境（相当于登录）（建议用这个）（退出用户：exit）
- **su 用户名** :切换到用户的身份（环境变量等没变，导致很多命令要加上绝对路径才能执行）
- **sudo 命令** ：以root的身份执行命令（输入用户自己的密码，而su为输入要切换用户的密码，普通用户需设置/etc/sudoers才可用sudo）

## 磁盘管理

- **df -h** :显示磁盘的空间使用情况 及挂载点
- **df -h /var/log** :（显示log所在分区（挂载点）、目录所在磁盘及可用的磁盘容量）
- **du -sm /var/log/\* | sort -rn** : 根据占用磁盘空间大小排序（MB）某目录下文件和目录大小
- **fdisk -l** :查所有分区及总容量，加/dev/sda为查硬盘a的分区）
- **fdisk /dev/sdb** :对硬盘sdb进行分区
- **mount /dev/sda1 /mnt** ：硬盘sda1挂载到/mnt目录（mount 装置文件名 挂载点）
- **mount -t cifs -o username=luolanguo,password=win用户账号密码,vers=3.0 //10.2.1.178/G /mnt/usb** :远程linux 共享挂载windows的U盘,G为U盘共享名，需设置U盘共享
- **mount -o loop /opt/soft/CentOS-7-x86_64-DVD-1708.iso /media/CentOS** ：挂载iso文件
- **umount /dev/sda1** ：取消挂载（umount 装置文件名或挂载点）

## 压缩、解压和打包备份

>  单纯tar仅为打包（多个文件包成一个大文件），加上参数-j(bzip2格式.bz2)、-z（gzip格式.gz）可以备份、压缩(-c)、解压（-x），备份一般比压缩多加参数-p（保留原本文件的权限与属性），-C可以指定解压到特定目录；bzip2、gzip只能对单一文件压缩；

- **file 文件名** ：查文件类型（可看是用哪一种方式压缩的）
- **tar -zxvf a.tar.gz -C   ./test** ：解压tar.gz到当前目录下的test目录
- **tar -zcvf /opt/c.tar.gz   ./a/** ：压缩tar.gz（把当前目录下的a目录及目录下所有文件压缩为 /opt/目录下的c.tar.gz）
- **tar -jxvf a.tar.bz2** ：解压tar.bz2（到当前目录）
- **tar -jcvf c.tar.bz2 ./a/** ：压缩tar.bz2（把当前目录下的a目录及目录下所有文件压缩到当前目录下为c.tar.gz2）
- **unzip a.zip** ：解压zip（到当前目录）
- **zip -r c.zip ./a/** :压缩zip(把当前目录下的a目录及目录下所有文件压缩到当前目录下为c.zip
- **bzip2 -k file1** ： 压缩一个 'file1' 的文件（-k表示保留源文件）（bzip2格式，比gzip好）
- **bzip2 -d -k file1.bz2** ： 解压一个叫做 'file1.bz2'的文件
- **gzip file1** ： 压缩一个叫做 'file1'的文件（gzip格式）（不能保留源文件）
- **gzip -9 file1** ： 最大程度压缩
- **gzip -d file1.gz** ： 解压缩一个叫做 'file1'的文件

## 软件安装

> - 尽量用yum源（apt-get）安装，不行就rpm、deb包安装，能不手动编译的就不要手动编译；
> - dpkg只能安装已经下载到本地机器上的deb包. apt-get能在线下载并安装deb包,能更新系统,且还能自动处理包与包之间的依赖问题,这个是dpkg工具所不具备的；
> - rpm 只能安装已经下载到本地机器上的rpm 包. yum能在线下载并安装rpm包,能更新系统,且还能自动处理包与包之间的依赖问题,这个是rpm 工具所不具备的;
> - yum、rpm安装文件分布在/usr的bin、lib、share不同目录，不用配置PATH，直接用命令，但可用命令卸载更新；
> - 手动编译软件，默认位置为/usr/local下不同子目录下,不用配置PATH直接用命令（手动指定安装路径需要加PATH），使得软件更新和删除变得很麻烦。编译安装的软件没有卸载命令，卸载就是把所有这个软件的文件删除。

### 二进制(Binaries)包

#### yum安装

>  在线下载并安装rpm包，适用于CentOS、Fedora、RedHat及类似系统

- **yum install epel-releas** ：安装第三方yum源EPEL（企业版 Linux 附加软件包的简称）
- **yum repolist enabled** ：显示可用的源仓库（/etc/yum.repos.d/目录下配置）
- **yum install yum-fastestmirror** ：自动选择最快的yum源
- **yum list installed |grep java** ：列出已安装的软件（查看已安装的JDK）
- **yum remove java-1.8.0-openjdk.x86_64** ：卸载软件（卸载JDK）
- **yum list java\*** ：列出已安装和可安装的软件（查看yum库中的JDK包）
- **yum install [-y] java-1.8.0-openjdk** ：安装软件JDK(-y自动安装)（推荐这种方式安装）
- **yum check-update [kernel]** ：列出所有可更新的软件（检查更新kernel）
- **yum update tomcat** ：更新软件（可所有）
- **rpm -ql 软件名称** ：查询yum安装路径（软件名称可通过rpm -qa|grep java）
- **yum info kernel** ：查看软件（kernel）的信息
- **yum clean all** ：（清除缓存，使最新的yum配置生效）

#### rpm包手动下载安装

>  yum中没有时用，适用于CentOS、Fedora、RedHat及类似系统；

- **wget -P /opt https://网址** ：下载到/opt目录
- **rpm -ivh wps-office-版本.x86_64.rpm** :安装rpm包（包要先下载）（要先装依赖包）
- **rpm -e wps-office** ：卸载软件（注意不要软件名不要版本号）
- **rpm -qa |grep wps** ：查看安装的rpm包
- **rpm -ql 软件名称** ：查看rpm包安装路径（软件名称可通过rpm -qa|grep java）

#### apt方式安装

>  安装deb包，类似yum安装，适用于Debian, Ubuntu 以及类似系统；

- **apt-get install aptitude** ：安装aptitude工具,实现依赖自动安装，依赖版本自动降级或升级
- **aptitude install 软件** ：安装软件（推荐这种方式安装）
- **apt-cache search 软件** ：搜索软件
- **apt-get install 软件** ：安装软件
- **apt-get purge 软件** ：卸载软件（包括配置文件，只删除软件purge换成remove）
- **apt-get upgrade** ：更新所有已安装的软件包
- **apt-get update** ：升级列表中的软件包
- **apt-get clean** ：从下载的软件包中清理缓存

#### deb包安装

>  适用于Debian, Ubuntu 以及类似系统；

- **dpkg -i package.deb** ：安装一个 deb 包
- **dpkg -r package_name** ：从系统删除一个 deb 包
- **dpkg -l |grep chrome** ：查询系统中所有已经安装的 deb 包
- **dpkg -L 软件名称** ：查软件安装的文件

#### 解压即用

>  大多数非开源的商业软件都采取这种办法；

 二进制（Binaries）包如[apache-jmeter-3.3.tgz](https://archive.apache.org/dist/jmeter/binaries/apache-jmeter-3.3.tgz)，下载复制解压到/opt，然后然后将该软件的 bin 目录加入到 PATH 中即可（vim /etc/profile export PATH=$PATH:/opt/apache-jmeter-3.3/bin）；

#### 软件自己的模块/包管理器

>  如python：系统的源中不可能包含该软件的所有模块； 系统的源中该软件的模块的更新要远远滞后于最新版本；手动安装python，并用Python 自带的 pip 安装模块（类似yum）；

- **pip install redis** ：安装python软件包[redis](http://www.ttlsa.com/redis/)
- **pip unstall redis** :卸载
- **pip show --files redis** :pip查看已安装的包
- **pip list --outdated** :检查更新

### 源代码(Source)包

#### 编译安装

>  源代码包（一般有install文件）如[hello-2.2.tar.bz2](http://ftp.gnu.org/gnu/hello/hello-2.2.tar.bz2)，下载复制到/opt;

- **tar -jxvf hello-2.2.tar.bz2** :解压
- **./configure --prefix=/opt/软件目录名称** :为编译做好准备，加上 prefix 手动指定安装路径
- **make** ：编译
- **make install** ：安装
- **make clean** ：删除安装时产生的临时文件
- **vim /etc/profile export PATH=$PATH:/opt/目录/bin** ：手动指定安装路径需要加path
- **hello** ：执行软件：看INSTALL和README文件（是否源码包、如何安装、执行都看这两个）
- **rm -rf 软件目录名称** :卸载软件

**转载请注明出处：https://www.cnblogs.com/caozy/p/9261224.html**