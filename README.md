# oak
基于ssh的小工具

# usage

```
# machine operate
oak m add 
oak m ls
oak m rm

# machine group
oak group add
oak group ls
oak group rm

# remote operate
oak exec [machine/group] [command]
oak ssh/con machine 
# copy file
oak cp [machine/group]/filepath  [filepath/path]
oak cp [file]/filepath [machine/group]/[filepath/path]
# copy dir
oak cp -r [machine/group]/path  [path]
oak cp [file]/path [machine/group]/path
```

## data Demo

```json
groups:
- machine: 
  - node1
  - node2
  name: g1

machines:
- name: node1
  host: 192.168.1.2
  password: 123456
  port: 22
  private_key: ""
  public_key: ""
  type: password
  user: root
- name: node2
  host: 192.168.1.2
  password: 123456
  port: 22
  private_key: ""
  public_key: ""
  type: password
  user: root
- name: node3
  host: 192.168.1.2
  password: 123456
  port: 22
  private_key: ""
  public_key: ""
  type: password
  user: root

```