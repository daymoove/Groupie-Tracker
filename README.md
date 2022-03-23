# Groupie tracker

## Sommaire

- [Lancement](#p1)
- [Différente fonctionnalité](#p2)
    - [Barre de recherche](#p2.1)
    - [Filtres](#p2.2)

## Lancement <a name="p1"></a>

Pour lancer notre projet il faut se mettre dans le dossier "server" puis run :
```
cd server
go run .
``` 

## Différente fonctionnalité <a name="p2"></a>

### Barre de recherche <a name="p2.1"></a>

Nous avons fais une barre de recherche qui fonctionne en live permettant de faire nos recherche sans que la page ne se rafraichisse.  
Elle permet de rechercher : 
- par le nom du groupe,  
- celui d'un membre,  
- leur date de création  
- la date de sortie de leur premier album.  

Elle se trouve en haut à droite de la page dans la barre bleue.  

### Filtres <a name="p2.2"></a>

Nous avons également mis en place des filtres permettant de faire apparaitre seulement les groupes qui sont contenue dans les filtres.  
On peut filtrer selon : 
- la localisation de leur concert,  
- leur date de création,  
- la date de sortie de leur premier album,  
- leur nombre de membre.  

On peut également les activer en même temps pour des recherches encore plus précise.  
Ils se trouve en haut à gauche de la page dans la barre bleue.  
