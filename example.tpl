<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Portfoleon Report</title>
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js" type="text/javascript"></script>
		<link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.11.4/css/jquery.dataTables.css">
		<script type="text/javascript" charset="utf8" src="https://cdn.datatables.net/1.11.4/js/jquery.dataTables.min.js"></script>
		<script>
		$(document).ready( function () {
			$('#myTable').DataTable();
		} );
		</script>
	</head>
	<body>
		<table id="myTable" class="table table-bordered table-hover table-condensed">
		<thead><tr><th title="name">name</th><th title="type">type</th><th title="status">status</th>{{range $key, $value := (index .data 0).fields }}<th title="{{ $key }}">{{ $key }}</th>{{end}}</tr></thead>
		<tbody>
		{{range .data }}<tr><td>{{.name}}</td><td>{{.work_item_type}}</td><td>{{.status}}</td>{{range  $key, $value := .fields }}<td>{{ $value }}</td>{{end}}</tr>
		{{end}}
		</tbody>
		</table>
	</body>
</html>