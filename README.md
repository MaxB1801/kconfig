# kconfig
kconfig

# resources used
https://stackoverflow.com/questions/53256373/sending-file-over-ssh-in-go

# Requirements
Windows
Linux

# Description
This code will automatically retrieve the kubeconfig from a given node and given cluster type. It can store the config in multiple filepaths
Use kconfig -edit to edit the config.yaml which configures the code

# New features
Supports Linux with nano

# Set Up Windows
Copy content into local directory, and add the directory to your path    
Edit contents of kconfig.yaml to your desired state    

# Set Up Linux
Copy contents into local directory  
Use command 'nano ~/.bashrc'  
Append to the end of the file "alias kconfig='<path_to_kconfig_folder>/kconfig'  
Edit contents of kconfig.yaml to your desired state 
