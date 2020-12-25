<?php

$pageFileName = "index.html";
$page = fopen($pageFileName, "r");
$pageFileSize = filesize($pageFileName);

echo fread($page, $pageFileSize);