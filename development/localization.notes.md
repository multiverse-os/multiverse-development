# Multiverse OS: Localization
===============================================================================
Multiverse OS design specifically specifies building out the functionality and
the tools to start providing translations for every string in every part of the
Multiverse OS code; from shell script help text, to installers, and all
aspects of the UI, including the https://multiverse-os.org. 

**Implementation**
Planning how this functionality can be implemented from the earliest pre-alpha
stages is very important, so that we can ensure we are building the logic 
and database functionality required, so YAML translation data can be loaded
from shell scripts to Go software. Doing this correctly will enable us to
the same base translation files for all Multiverse OS software, requiring only
minor additions to the base to complete translation for each piece of software.

**Functional Requirements**

  * Implement loading YAML data in Shell, Go, Rust, and C. It may be important
    to load it in Bash too. 

    The easiest way may be using a compied Go program to load YAML files into
    space separated strings, in groupings defined by sections and subections. 

  * Ability to move around key words around for translations, the variables 
    in a translation must be carried into the trnaslation, because word 
    order is dependent on language.  
