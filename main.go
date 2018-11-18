package main

import (
	"clientWidget2/initClient"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type ConnInfo struct {
	pass string
	user string
	port int
}

var loginForm *widgets.QWidget

func main() {

	widgets.NewQApplication(len(os.Args), os.Args)

	loginForm = LoginForm()
	loginForm.Show()

	widgets.QApplication_Exec()
}

func LoginForm() *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/login.ui")

	file.Open(core.QIODevice__ReadOnly)
	var formWidget = loader.Load(file, widget)
	file.Close()

	var (
		passwordEdit = widgets.NewQLineEditFromPointer(widget.FindChild("passwordEdit", core.Qt__FindChildrenRecursively).Pointer())
		usernameEdit = widgets.NewQLineEditFromPointer(widget.FindChild("usernameEdit", core.Qt__FindChildrenRecursively).Pointer())
		portEdit     = widgets.NewQLineEditFromPointer(widget.FindChild("portEdit", core.Qt__FindChildrenRecursively).Pointer())
		okButton     = widgets.NewQPushButtonFromPointer(widget.FindChild("okButton", core.Qt__FindChildrenRecursively).Pointer())
		//exitButton   = widgets.NewQPushButtonFromPointer(widget.FindChild("exitButton", core.Qt__FindChildrenRecursively).Pointer())
		loginRadio = widgets.NewQRadioButtonFromPointer(widget.FindChild("loginRadio", core.Qt__FindChildrenRecursively).Pointer())
		//registrationRadio = widgets.NewQRadioButtonFromPointer(widget.FindChild("registrationRadio", core.Qt__FindChildrenRecursively).Pointer())
	)

	okButton.ConnectClicked(func(a bool) {
		pass := passwordEdit.Text()
		user := usernameEdit.Text()
		port := portEdit.Text()
		var chat initClient.ChatClient
		chat.Init(port)
		if loginRadio.IsChecked() { //so it's a login then

			if chat.Login(user, pass) {
				box := widgets.NewQMessageBox(nil)
				box.SetText("Log in was suckassful")
				box.Exec()
				chat := ChatForm(&chat)
				chat.Show()
				loginForm.Hide()
			} else {
				box := widgets.NewQMessageBox(nil)
				box.SetText("Log in was unsuckassful")
				box.Exec()
			}
		} else {
			if chat.Registration(user, pass) {
				box := widgets.NewQMessageBox(nil)
				box.SetText("Registration was suckassful")
				box.Exec()
			} else {
				box := widgets.NewQMessageBox(nil)
				box.SetText("Registration was unsuckassful")
				box.Exec()
			}
		}
	})
	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(formWidget, 0, 0)
	widget.SetLayout(layout)

	widget.SetWindowTitle("Chat client")

	return widget
}

func ChatForm(chat *initClient.ChatClient) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/form.ui")

	file.Open(core.QIODevice__ReadOnly)
	var formWidget = loader.Load(file, widget)
	file.Close()

	var (
		textEdit     = widgets.NewQTextEditFromPointer(widget.FindChild("textEdit", core.Qt__FindChildrenRecursively).Pointer())
		button       = widgets.NewQPushButtonFromPointer(widget.FindChild("pushButton", core.Qt__FindChildrenRecursively).Pointer())
		chatTextEdit = widgets.NewQTextEditFromPointer(widget.FindChild("textEdit_2", core.Qt__FindChildrenRecursively).Pointer())
	)
	go initClient.WaitForResponse(chat.Server, chatTextEdit)
	button.ConnectClicked(func(a bool) {
		chat.SendMessage(textEdit.ToPlainText())
		textEdit.Clear()
	})
	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(formWidget, 0, 0)
	widget.SetLayout(layout)

	widget.SetWindowTitle("Chat client")

	return widget
}
