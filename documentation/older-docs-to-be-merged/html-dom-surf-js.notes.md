# Manipulating DOM, executing JS and improving the Surf Go Library

* Need more than just Body() and HTML(), because sometimes we will be downloading Javascripts and Stylesheets

* Browser.Open() should be able to guess a scheme if its not present, its a browser. Start with 443, then 80, then quit.

## JS Execution

var re=new RegExp('/[a-z]{2}_[A-Z]{2}/','g');
var url=window.location.href;
var firstIndex=url.search(re);
var lang="de_DE";

switch (lang) {
  case "en_GB": lang="English"; break;
  case "de_DE": lang="Deutsch"; break;
  case "it_IT": lang="Italia"; break;
  case "es_ES": lang="Spanish"; break;  
}

var newOpt1 = new Option("Language: "+lang, "", "true", "trueS");
var selector = document.getElementById('language_selector');
selector.options[0]=newOpt1;
jQuery('#language_selector, #language_selector_desktop').change(function(){
  /**/if (jQuery(this).val() != '0') window.location.replace('/change-language?language='+jQuery(this).val());/**/
})

var _switchBackground=6000;
function chgBackgrnd(o)
{
    jQuery('.Tstgco > img > .Takt')
    if (typeof(o) == 'undefined' || !o) {
      var o = jQuery('.Tstgrtr s.Takt').next();
    } 
    if (!o.length) { o = jQuery('.Tstgrtr s').first(); }
    if (o) {
      if (!o.hasClass('Takt')){
         changeToStage(o.attr('num'));
      }
      setTimeout('chgBackgrnd()',_switchBackground);

      if(TisMOB && jQuery('#Tstage > img').css('display') != 'none' ){  
         jQuery('#Tstglogbox').first().css('margin-top','180px'); } 
    } else {
    }

    if(TisMOB){
      jQuery('.Tstgco > img').first().attr('src','/shared/static_fon/imgs/mob_keyvisual_1,tid=da.png');
      jQuery('.Tstgco > img').last().attr('src','/shared/static_fon/imgs/mob_keyvisual_2,tid=da.png');
    }
}
/*setTimeout('chgBackgrnd()',_switchBackground);*/

[JSVM] Output undefined
[ERROR] TypeError: Cannot access member 'href' of undefined

[BROWSER:PARSING] len(tag.Text()) is 1508
[BROWSER:PARSING] tag.Text() is 
var re=new RegExp('/[a-z]{2}_[A-Z]{2}/','g');
var url=window.location.href;
var firstIndex=url.search(re);
var lang="de_DE";

switch (lang) {
  case "en_GB": lang="English"; break;
  case "de_DE": lang="Deutsch"; break;
  case "it_IT": lang="Italia"; break;
  case "es_ES": lang="Spanish"; break;  
}

var newOpt1 = new Option("Language: "+lang, "", "true", "trueS");
var selector = document.getElementById('language_selector');
selector.options[0]=newOpt1;
jQuery('#language_selector, #language_selector_desktop').change(function(){
  /**/if (jQuery(this).val() != '0') window.location.replace('/change-language?language='+jQuery(this).val());/**/
})

var _switchBackground=6000;
function chgBackgrnd(o)
{
    jQuery('.Tstgco > img > .Takt')
    if (typeof(o) == 'undefined' || !o) {
      var o = jQuery('.Tstgrtr s.Takt').next();
    } 
    if (!o.length) { o = jQuery('.Tstgrtr s').first(); }
    if (o) {
      if (!o.hasClass('Takt')){
         changeToStage(o.attr('num'));
      }
      setTimeout('chgBackgrnd()',_switchBackground);

      if(TisMOB && jQuery('#Tstage > img').css('display') != 'none' ){  
         jQuery('#Tstglogbox').first().css('margin-top','180px'); } 
    } else {
    }

    if(TisMOB){
      jQuery('.Tstgco > img').first().attr('src','/shared/static_fon/imgs/mob_keyvisual_1,tid=da.png');
      jQuery('.Tstgco > img').last().attr('src','/shared/static_fon/imgs/mob_keyvisual_2,tid=da.png');
    }
}
/*setTimeout('chgBackgrnd()',_switchBackground);*/

[JSVM] Output undefined
[ERROR] TypeError: Cannot access member 'href' of undefined


## DOM




## HTML

https://github.com/xyproto/onthefly
