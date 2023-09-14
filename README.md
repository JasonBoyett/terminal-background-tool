# Terminal Background Selector

## Description
This is a terminal application to help users choose their terminal background 
image from the command line.
Though this application was built for choosing terminal backgrounds,
it can be used for any application that renders its background from an image file at a
given path.


## Validated Platforms
- Kitty
- Warp

(If you have tested this application on another platform,
please submit a pull request to update this list and document the setup process.)

## Installation
1. Make sure you have [Go](https://go.dev/doc/install) installed on your machine.
2. Make sure you have [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
3. Clone this repository
    ```bash
    git clone https://github.com/JasonBoyett/terminal-background-tool
    ```
4. Navigate to the repository
    ```bash
    cd terminal-background-tool
    ```
5. Build the application
    ```bash
    go build cmd/main/main.go
    ```
6. Run the application
    ```bash
    ./main
    ```

## Usage
The first time you run the application it will ask you what directory you would like
to use as your background image directory. If this directory does not exist, it will 
automatically be created. If you would like to change this directory, you can do so in the 
config.json file by changing the "bgDirectory" field.

On initial setup, the application will create an "images" directory in the directory you
specified. This is where you will place your background images. There will be two other directories labeled "png_images" and "jpg_images" These will store the current background image.

After initial setup, you can run the application again and will be prompted to choose 
from a list of background images. Navigate with your arrow keys or with the j and k keys.
Press enter to choose an image, r to select one randomly, or q to quit.
If you would like to add more images, simply place them
in the "images" directory and run the application again.

## Setup for Kitty
In your kitty.conf file add the following lines:
```
background_image [your backgrounds directory]/png_images/current.png
background_image_layout scaled
```
you will have to replace [your backgrounds directory] with the directory you specified
when you first ran the application. I would recommend copying and pasting it from the
config.json file.

Now after running the application reload kitty and you should see your new background.


