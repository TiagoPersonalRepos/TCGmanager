# Development Notes #

## Installation of **GO** and **Fyne** ##
~~~ Bash
sudo apt install golang gcc libgl1-mesa-dev xorg-dev # Assuming it's a Debian based OS
go get fyne.io/fyne/v2@latest
echo "export PATH=$PATH/\"$HOME/go/bin\"" >> ~/.bashrc && source ~/.bashrc
~~~

## Testing **GO** and **Fyne** ##
~~~ Bash
go version
fyne --version
~~~

## Testing Fyne Demo ##
~~~ Bash
go run fyne.io/fyne/v2/cmd/fyne_demo@latest
~~~

## Installing Android Tools ##
Without isntalling Android-Studio

~~~ Bash
sudo apt install adb sdkmanager # Assuming it's a Debian based OS
sdkmanager --version # use 'sudo' if it doesn't work
sdkmanager "platform-tools" "platforms;android-33"
sdkmanager --install "cmake;3.22.1" "ndk-bundle;25.0.8775105" "build-tools;33.0.1" "tools"  "build-tools;33.0.0" "emulator;33.1.1"
~~~

fyne package -os android -appID com.example.myapp