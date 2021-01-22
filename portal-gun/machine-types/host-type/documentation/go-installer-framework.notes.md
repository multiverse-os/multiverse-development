## 'vfio-bind' script Go conversion notes
Below are a list of notes taken while doing the latest rewrite of the 'vfio-bind' script and associated 'vfio' module for the shell framework that accompanies Multiverse OS. 

	* Don't just simply list the PCI addresses found for a given device ID. But go before anything, parse the entire list of P:CI devices, organize by category (type, and various other groupings, likely using a basic tag system for filtering.) Then in the CLI software, have a list that one can go up/down, get more information on any item that unfolds additional data (with only one showing at a time) and finally multi-select and submission for binding to 'vfio' for later PCI passthrough

