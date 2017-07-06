# NSQ prezentacija

Prezentacija opisuje kako radi i kako se koristi [NSQ](http://nsq.io/), te kako ga koristimo u [Minus5](http://www.minus5.hr/).

## Pokretanje prezentacije

Prije pokretanja, potrebno je imati instaliran [Consul](https://www.consul.io/) i [NSQ](http://nsq.io/)
Potrebni su zato što prezentacija za NSQ primjer koristi Consul service discovery o kojima ovise naše [svckit biblioteke](https://github.com/minus5/svckit).

[Consul](https://www.consul.io/) i [NSQ](http://nsq.io/) je na MAC-osu najlakše instalirati kroz [Homebrew](https://brew.sh/).

```
# brew install consul
# brew install nsq
```

Prezentacija se pokreće shell skriptom:
```
# ./test.sh
```

Nakom pokretanja, prezentacija je dostupna na:  
[http://localhost:8080](http://localhost:8080)

Za one koje zanima samo "Super Mario" mogu ga pogledati na:  
[http://localhost:8080/mario_demo.html](http://localhost:8080/mario_demo.html)

## Struktura foldera
./prez sadrži prezentaciju koja je napravljena sa:  
[http://lab.hakim.se/reveal-js/](http://lab.hakim.se/reveal-js/)

consul.json je statička consul konfiguracija kako bi svckit mogao korisiti nsq

main.go je primjer za nsq komunikaciju a ujedno i server

./prez/js/nsqdemo.js je JS za komunikaciju WS->NSQ->WS kao dio primjera.

draw.io sadrži slike napravljene sa:  
[https://www.draw.io/](https://www.draw.io/)

ostala grafika u zvukovi su pokupljeni s interneta..
- google "super mario"
- google "super mario sprites"
- google "super mario sounds"

## Reference
- NSQ [http://nsq.io/](http://nsq.io/)
- Consul [https://www.consul.io/](https://www.consul.io/)
- Minus5 svckit [https://github.com/minus5/svckit](https://github.com/minus5/svckit), [primjeri](https://github.com/minus5/svckit/tree/master/example)
- Reveal.js [http://lab.hakim.se/reveal-js/](http://lab.hakim.se/reveal-js/)
- Super Mario [http://mario.nintendo.com/](http://mario.nintendo.com/)
- Super Mario Wii editor [http://rvlution.net/reggie/](http://rvlution.net/reggie/)
- Super Mario zvukovi [http://www.mariomayhem.com/downloads/sounds/](http://www.mariomayhem.com/downloads/sounds/)