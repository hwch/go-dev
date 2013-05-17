package main

import (
        "gtk"
        "os"
        "unsafe"
)

func print_hello() {
        println("Button1")
}

func main() {
        gtk.Init(os.Args)
        win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
        win.Connect("destroy", func() {
                gtk.MainQuit()
        })
        win.SetTitle("Hello World")
        cn := gtk.ToContainer(unsafe.Pointer(win.CType()))
        cn.SetBorderWidth(10)
        // cn.SetResizeMode(gtk.RESIZE_PARENT)
        gr := gtk.NewGrid()

        cn.Add(gtk.ToWidget(unsafe.Pointer(gr.CType())))
        bt := gtk.NewButtonWithLabel("Button")
        bt.Connect("clicked", func() {
                println("Hello World")
        })
        gr.Attach(gtk.ToWidget(unsafe.Pointer(bt.CType())),
                0, 0, 1, 1)
        bt1 := gtk.NewButtonWithLabel("Button1")
        bt1.Connect("clicked", print_hello)
        gr.Attach(gtk.ToWidget(unsafe.Pointer(bt1.CType())),
                1, 0, 1, 1)
        bt2 := gtk.NewButtonWithLabel("Quit")
        bt2.Connect("clicked", gtk.MainQuit)
        gr.Attach(gtk.ToWidget(unsafe.Pointer(bt2.CType())),
                0, 1, 2, 1)
        gr.SetColumnHomogeneous(true)
        gr.SetRowHomogeneous(true)
        gr.SetColumnSpacing(10)
        gr.SetRowSpacing(5)
        win.ShowAll()
        gtk.Main()
}
