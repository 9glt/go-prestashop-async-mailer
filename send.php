<?php

$postdata = http_build_query(
    array(
        'to' => 'some@email.tld',
        'subject' => 'some subject',
    )
);

$opts = array('http' =>
    array(
        'method'  => 'POST',
        'header'  => 'Content-Type: application/x-www-form-urlencoded',
        'content' => "this is email content as html"
    )
);

$context  = stream_context_create($opts);


$result = file_get_contents('http://localhost:8888/send?'.$postdata, false, $context);
