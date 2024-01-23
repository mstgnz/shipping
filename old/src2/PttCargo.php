<?php

class PttCargo extends Cargo
{

    public function __construct($username, $password, $urlMode = "test", $urls = array())
    {
        parent::__construct($username, $password, $urlMode, $urls);
    }

    public function createToCargo($query = array())
    {
        return "ptt ile kargo oluştur";
    }

    public function whereIsTheCargo($query = array())
    {
        return "ptt ile kargo sorgula";
    }

    public function cancelToCargo($query = array())
    {
        return "ptt ile kargo iptal";
    }
}
