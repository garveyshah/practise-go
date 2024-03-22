#! /bin/bash

curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq -c ' .[] | select ( .id == '$HERO_ID') | .connections.relatives' | tr -d '"'
