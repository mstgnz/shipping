<?php

abstract class Cargo
{
    protected $soap;
    protected $username;
    protected $password;
    protected $urlMode;
    protected $urls;

    // Tüm kargo entegrasyonlarında standart olan veriler
    protected function __construct($username, $password, $urlMode = "test", $urls = array())
    {
        $this->username = $username;
        $this->password = $password;
        $this->urlMode = $urlMode;
        $this->urls = $urls;
        $this->initialize();
    }

    private function initialize()
    {
        try {
            $this->soap = new \SoapClient($this->urls[$this->urlMode], array('trace' => true));
        } catch (\Throwable $th) {
            return $th->getMessage();
        }
    }

    // Tüm kargo entegrasyonlarında standart olan (oluşturma - sorgulama - iptal) işlemleri -> her kargo firması kendine göre yazacak
    protected abstract function createToCargo($query = array());
    protected abstract function whereIsTheCargo($query = array());
    protected abstract function cancelToCargo($query = array());

    // Array check for key name
    protected function arrayCheck($args, $key)
    {
        $result = false;
        if (is_array($args)) {
            foreach ($this->necessary[$key] as $value) {
                $result = array_key_exists($value, $args);
            }
        }
        return $result;
    }

    // Xml Output
    protected function xmlOutput($array)
    {
        $xml_data = new SimpleXMLElement('<?xml version="1.0"?><data></data>');
        $this->array_to_xml($array, $xml_data);
        header('Content-Type: application/xml; charset=utf-8');
        echo $xml_data->asXML();
        exit;
    }

    // Soap Get Last Request
    protected function getLastRequest()
    {
        echo $this->soap->__getLastRequest();
        exit;
    }

    // Arra Convert To XML
    protected function array_to_xml($data, &$xml_data)
    {
        foreach ($data as $key => $value) {
            if (is_array($value)) {
                if (is_numeric($key)) {
                    $key = 'item' . $key;
                }
                $subnode = $xml_data->addChild($key);
                $this->array_to_xml($value, $subnode);
            } else {
                $xml_data->addChild("$key", htmlspecialchars("$value"));
            }
        }
    }
}
