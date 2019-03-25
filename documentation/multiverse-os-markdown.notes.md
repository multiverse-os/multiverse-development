# Multiverse OS: Markdown
===============================================================================
Multiverse OS requires markdown features that do not yet exist in the current
flavors of Multiverse. In order to provide the functionality needed to meet the
design requirements of the multiverse-os.org community site, documentation, 
peer review community help documents, and simple GUI implementation using
webframe with webkit and markdown; enabling even designers with no desktop
experience but basic markdown (or HTML and CSS) knowledge could rapidly
prototype and test desktop GUI interfaces using markdown. 


===============================================================================
### Markdown
Functional requirements for Multiverse flavored markdown:

 * **Reference: local or remote file content inclusion**
   Ability to reference a local file or remote document, or git repository. 
   Specifying the path (remote or local), the starting line, and the
   last line, copies the content into a quote, a code block (with language
   explicitly specified, or other details about the source.
    
   **Example usage** developers can decide to refernce a section of source
   code in the documentation using this feature, and any updates to the source
   code will automatically be reflected in the docuemntation without any manual
   updates needed, and preventing any confusion due to the documentation source
   being out of sync with the actual source. This provides higher quality
   documentation (and more) with less work. 

 * **Grapviz: a codeblock of graphviz code can include an additional component**
   in the declaration of the codeblock that renders the graphviz image and
   includes it inline using dataurl+base64. 

 * **All images are converted to base64 and included using data URLs**

 * **Ability to render charts/graphs from a table of data with an adittional**
   component in the delcaration of the table of data. This graph is rendered
   preferably as SVG+CSS and JS if necessasry for antimation. 

   Ability to use remote data source to feed live data into graph. 

 


