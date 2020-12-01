{{-- resources/views/admin/dashboard.blade.php --}}

@extends('adminlte::page')

@section('title', 'Dashboard')

@section('content_header')
@stop

@section('content')
    <div class="container-fluid">
        <div class="row">
            <div class="col-md-12">
                <div class="card">
                    <div class="card-header">
                        <h3 class="card-title">Список активных клиентов</h3>

                        <div class="card-tools">
                            <div class="input-group input-group-sm" style="width: 450px;">
                                <input type="text" id="search" name="table_search" class="form-control float-right" placeholder="Поиск">

                                <div class="input-group-append">
                                    <button type="submit" class="btn btn-default"><i class="fas fa-search"></i></button>
                                </div>
                            </div>
                        </div>
                    </div>
                    <!-- /.card-header -->
                    <div class="card-body p-0">
                        <table class="table table-condensed">
                            <thead>
                            <tr>
                                <th>Клиент</th>
                                <th>Программа</th>
                                <th>Группа</th>
                            </tr>
                            </thead>
                            <tbody>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
@stop

@section('js')
    <script>

        function delay(fn, ms) {
            let timer = 0
            return function(...args) {
                clearTimeout(timer)
                timer = setTimeout(fn.bind(this, ...args), ms || 0)
            }
        }

        function search(query) {
            $.ajax({
                dataType: "json",
                url: '/api/clients/search',
                data: {query: query},
                success: function(data) {

                    $('tr.dynamic').remove();

                    $.each( data.data, function( key, val ) {
                        $("tbody")
                            .append($('<tr class="dynamic">')
                                .append($('<td>')
                                    .append(val.name)
                                )
                                .append($('<td>')
                                    .append(val.program.name)
                                )
                                .append($('<td>')
                                    .append(val.group.name)
                                )
                            );
                    });
                }
            });
        }

        search();

        $('body').on('keyup', '#search', delay(function() {
            search($(this).val());
        }, 500))
    </script>
@stop
