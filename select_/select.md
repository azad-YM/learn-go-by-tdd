un `select` est un mécanisme utilisé pour écouter plusieurs canaux (channels) en même temps. 
Il permet de gérer la concurrence proprement, un peu comme un switch mais pour les opérations sur channels.

Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool. Since we are closing and not sending anything on the chan, why allocate anything?

Always make channels

Notice how we have to use make when creating a channel; rather than say var ch chan struct{}. When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.

For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels


You'll recall from the concurrency chapter that you can wait for values to be sent to a channel with myVar := <-ch. This is a blocking call, as you're waiting for a value.

select allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.