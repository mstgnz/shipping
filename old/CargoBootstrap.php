<?php

function load_file(){
    $directory = __DIR__."/src/";
    $myfiles = array_diff(scandir($directory), array('.', '..')); 

    foreach ($myfiles as $key => $value) {
        if (file_exists($directory.$value)) {
            require_once $directory.$value;
        }
    }
	
}

load_file();

?>
