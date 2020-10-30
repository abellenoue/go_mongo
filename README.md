# go_mongo

# binome : 
    Vincent HOUDAN et Antoine BELLENOUE

# stack :
 *   Depot : github
 *   conteuneur : Docker
 *   cloud : AWS
 *   accées cloud : IAM
 *   IDE : VS-code
 *   pipeline : cloudShare
 *   DB : mongoDB atlas
 *   registre : ECR (elastique Conteuneur resitry)
 *   ec2 ou esc

# justification :

 1. GitHub : outil trés largement populaire. Bon nombre de logiciel et service peuvent se connecter.

 2. Docker : seul outil de contairisation vu en cours

 3. AWS : cloud avec la possibilité de migrée et/ou repliqué l'infrastructure partout dans le monde
    Etant trés connus, un bon nombre de tutorial est disponible et la docuentation est trés compléte

 4. IAM : le gestionnaire d'identification par default de aws, il nous permet de nous connecter partout

 5. VS-CODE : c'est gratuit, a l'option live-share pour travaillé a deux sur une même machine.
    vs-code a aussi un terminal intégrer et est cross-platform

 6. cloudshare : c'est gratuit, ça intégre tres bien aws et github. courbe d'apprentissage légére, facile a prendre en main, 

 7. mongoDbAtlas : permet une replication simple et est facile a gérer

 8. ECR : depot privé dans aws pour stocker les images docker, il est celuis qui offre le plus de sécurité

 9. EC2 ou ESC : deux solution de aws, une est du IASS et l'autre un PASS

# Acces
liens a l'API http://15.237.119.161:4747/