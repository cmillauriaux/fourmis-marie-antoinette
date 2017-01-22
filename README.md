# Le projet
Ce dépôt contient le code source du site [Marie-ANToinette](http://www.fourmis-marie-antoinette.fr), permettant d'observer l'évolution d'une colonie de fourmis en temps réel et de remonter son historique.

L'application est hebergée sur Google AppEngine et est un laboratoire pour la mise en place de différentes technologies : Go, Google Cloud, Angular 2, Ionic 2, etc.

Le projet sera régulièrement mis à jour pour intégrer de nouvelles fonctionnatlités ou expérimenter de nouveaux frameworks.

## Arborescence
Le dépôt est découpé en trois projets :
- *api* : Le serveur d'application exposant une API REST et permettant de servir l'application web
- *camera* : Une petite application Go et un script bash s'appuyant sur une Raspberry Cam pour prendre des clichés toutes les minutes et les envoyer sur Google Cloud
- *scripts* : Script de déploiement pour Google AppEngine
- *site* : Application web écrite en Angular2 pour offrir aux utilisateurs une manière simple et intuitive de naviguer dans l'historique de la fourmilière