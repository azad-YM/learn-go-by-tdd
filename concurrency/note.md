

Les maps en go, ne supportent pas que plusieurs éléments tentent d'y écrire simultanement.

Il s'agit d'une condition de concurrence, un bug qui survient lorsque le résultat de notre logiciel dépend du moment et de la séquence d'événements sur lesquels nous n'avons aucun contrôle. 

Comme nous ne pouvons pas contrôler précisément le moment où chaque goroutine écrit dans la table des résultats, nous sommes vulnérables à ce que deux goroutines y écrivent simultanément. 