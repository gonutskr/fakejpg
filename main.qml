import QtQuick 2.0
import QtQuick.Dialogs 1.1

Rectangle {
	id: root

	width: 640
	height: 480

	Rectangle {
			id: openBtn
			x: 300; y:0
			color: "gray"
			border.color: "black"
			width: 50
			height: 20

			Text {
				text: "..."
				anchors.centerIn: parent
				color: "black"

				MouseArea {
					anchors.fill: parent
					onClicked: { fileDialog.open() }
				}
			}
	}

	FileDialog {
		id: fileDialog
		selectFolder: true
		onAccepted: { 
			input.text = fileUrl 
			ctrl.onAcceptedBtnClicked(parent, fileUrl)
		}

	}

	TextInput {
		id: input
		x: 0; y: 0
		width: 300
		height: 20
		cursorVisible: false
	}

	Text {
		id: result
		objectName: "result"
		x: 0; y: 20
		width: 300
		height: 460
	}


}
