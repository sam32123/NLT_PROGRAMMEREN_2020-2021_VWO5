# Doelen en stappen
Hier zullen alle doelen en stappen beschreven staan. 

## Doel
Ons doel is een app die ervoor zorgt dat er minder voedselverspilling plaatsvindt door gemakkelijk recepten te kunnen zoeken met overgebleven etenswaren.


### Hoe gaan we dat bereiken?
* Recepten opslaan uit de Allerhande. 
* Producten vinden
    * Op basis van een barcode. 
    * Op basis van image recognition. 
    * Op basis van een zoekopdracht (naam). 
* In de recepten zoeken naar de producten die de gebruiker heeft ingevoerd. 

## Verschillende onderdelen om te maken

### app
* connectie met de backend
* resultaten van backend weergeven
* zoekfuncties
   * zoekfunctie op basis van text
   * zoekfunctie op basis van barcode
* preview van recepten, als je erop klinkt link ernaar

### backend
* requests kunnen sturen naar albert heijn site voor inredienten
* requests van de albert heijn site kunnen cashen
* requests kunnen sturen naar de allerhande site voor recepten op basis van ingredienten
* requests van de allerhande kunnen cashen
