> This project created for Beykent University's Computer Science Community.

Main part of this project is removing the background of an image and replacing it with different image while resizing and repositioning.

# Steps

### RabbitMQ
> This is the main part for communicating client and server

1. Run the Docker Desktop Client
   > You can use it from the terminal but I prefer this way to run docker images.
3. Run the RabbitMQ container command supplied from official RabbitMQ docs.
   ```
   docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.0-management
   ```


### Python Server
> This is the server that removes the background of the image
1. Open new terminal and navigate to project's directory
   > Or simply open the project's directory in an editor and open the terminal in there
2. Install the dependencies for server with included Makefile. Run this command
   ```
   make serverDep
   ```
3. Run the server with this command
   ```
   make server
   ```
Now the server is ready for requests

### Golang Client
> This is the client that processes images and places, resizes them.

The golang client accepts images from resources folder with proper naming and extension.
`image.jpg` for the image that you want to remove it's background and `background.jpg` for the image you want it to be the background.

> [!WARNING]
> The naming conventions is important. And just use __.jpg__ for input images. Or you can get errors while processing
 
1. Open another terminal and navigate to project's directory
   > Or simply in the editor, open a new terminal
2. Run the client with this command
   ```
   make run
   ```
Ta daa! you successfully generated an image!

> [!IMPORTANT]
Generated images saves to `resources/gen` directory. You can find it here with `generated.png` name.
