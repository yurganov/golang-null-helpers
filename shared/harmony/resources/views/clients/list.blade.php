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
                        <h3 class="card-title">Управление клиентами</h3>

                        <div class="card-tools">
                            <div class="input-group input-group-sm" style="width: 450px;">
                                <input type="text" name="table_search" class="form-control float-right" placeholder="Поиск">

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
                            @foreach($clients as $key => $client)
                                <tr>
                                    <td>{{ $client->client_name }}</td>
                                    <td>{{ $client->program_name }}</td>
                                    <td>{{ $client->group_name }}</td>
                                </tr>
                            @endforeach
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
@stop
