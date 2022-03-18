<h1 align="center">COMITEAME</h1>

#### [TO: English](#English)
#### [TO: Spanish](#Spanish)

# English

### What is comiteame?
My colleagues and I needed a way to do commits following our stablished semantic rules fast and intuitive.
Thats why I'm developing this little go program, using the libraries [promptui](https://github.com/manifoldco/promptui) and
[go-git](https://github.com/go-git/go-git).

### How to use it?
You can compile it by yourself if you have **go** in your computer. Just change dir to **cmd/comiteame/** and run **go build -o comiteame**

Also, you can download a compiled-ready binary from the releases tab.

You should deposit your binary to any directory in your system that is added to the PATH varible. This way you should be able to execute the
program in whatever git directory you are.

Just execute comiteame and the program will guide your througth the steps you need to follow.

# Spanish

### ¿Para que sirve esto?
Mis compañeros y yo necesitabamos una manera de hacer los commits siguiendo las reglas semánticas que establecimos rápida y intuitiva.
Para ello, me encuentro desarrollando este pequeño programa en go, que, usando las librerías [promptui](https://github.com/manifoldco/promptui) y 
[go-git](https://github.com/go-git/go-git), me permiten realizar este programa de manera muy sencilla y rápida.

### ¿Cómo se usa?
Puedes compilarlo por tu cuenta si tienes instalado **go** en tu ordenador. Simplemente dirígete al directorio **cd cmd/comiteame/**
y ejecuta **go build -o comiteame**

O bien, desde la pestaña de releases, puedes descargarte ya un binario compilado.

Es importante que el binario lo situes en cualquier parte de tu S.O que esté en la variable PATH del sistema. De esta manera podrás ejecutar 
el programa en cualquier directorio git de tu sistema.

Simplemente ejecuta **comiteame** y el propio programa te guiará con los pasos que tienes que seguir.
