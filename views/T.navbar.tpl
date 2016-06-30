{{define "t_navbar"}}
<nav class="navbar navbar-default navbar-fixed-top col-md-12 col-sm-12 col-xs-12">
        <div class="container-fluid">
            <!-- Brand and toggle get grouped for better mobile display -->
            <div class="navbar-header">
                <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1"
                    aria-expanded="false">
        <span class="sr-only">Toggle navigation</span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
                <a class="navbar-brand" href="#">WebNote</a>
            </div>

            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                <ul class="nav navbar-nav">
                    <li class=""><a href="#">List <span class="sr-only">(current)</span></a></li>
                    <li class="dropdown">
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Category <span class="caret"></span></a>
                        <ul class="dropdown-menu">
                            <li><a href="#">C_1</a></li>
                            <li><a href="#">C_2</a></li>
                            <li><a href="#">C_3</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="#">Separated link</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="#">One more separated link</a></li>
                        </ul>
                    </li>
                </ul>
                <form class="navbar-form navbar-left" role="search">
                    <div class="form-group">
                        <input type="text" class="form-control" placeholder="Search Note">
                    </div>
                    <button type="submit" class="btn btn-default">Search</button>
                </form>
                <!-- Button trigger modal -->
                <button type="button" class="btn btn-default  navbar-btn" data-toggle="modal" data-target="#addNoteModal">Add New Note</button>
                <!--登录注册-->
                <ul class="nav navbar-nav navbar-right {{if .isLogin}}hidden{{end}}">
                    <li><a href="javascript:void(0);" data-toggle="modal" data-target="#SignupModal">SignUp</a></li>
                    <li><a href="javascript:void(0);" data-toggle="modal" data-target="#loginModal">LogIn</a></li>
                </ul>
                <!--登录注册-->
                <!--个人信息-->
                <ul class="nav navbar-nav navbar-right {{if .isLogin}}{{else}}hidden{{end}}">
                    <li><a href="#"><img src="/static/upload/{{.info.Photo}}" alt="..." class="img-circle img-responsive" width="25px"></a></li>
                    <li class="dropdown">
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">{{.info.Name}} <span class="caret"></span></a>
                        <ul class="dropdown-menu">
                            <li><a href="#">Info</a></li>
                            <li><a href="#">Notes</a></li>
                            <li><a href="#">Setting</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="/admin">Logout</a></li>
                        </ul>
                    </li>
                <!--个人信息-->
                </ul>
            </div>
            <!-- /.navbar-collapse -->
        </div>
        <!-- /.container-fluid -->
    </nav>


    <!--模态框-->
    <!-- Add Note -->
    <div class="modal fade" id="addNoteModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="myModalLabel">Note Form</h4>
                </div>
                <div class="modal-body">
                    <form action="/note" method="POST">
                        <div class="form-group">
                            <label for="">Note Title</label>
                            <input type="text" class="form-control" id="ntitle" name="ntitle" placeholder="Title">
                        </div>
                        <div class="form-group">
                            <label for="">Note Category</label>
                            <input type="text" class="form-control" id="ntag" name="ntag" placeholder="category">
                        </div>
                        <div class="form-group">
                            <label for="">Note Contents</label>
                            <textarea class="form-control" rows="15" id="ncontent" name="ncontent" placeholder="Input you want"></textarea>
                        </div>                        
                        <button type="submit" class="btn btn-default btn-block">Submit</button>
                    </form>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                </div>
            </div>
        </div>
    </div>
    <!-- Login -->
    <div class="modal fade" id="loginModal" tabindex="-1" role="dialog" aria-labelledby="loginModalLabel">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="loginModalLabel">Note Form</h4>
                </div>
                <div class="modal-body">
                    <form method="POST" action="/admin">
                        <div class="form-group">
                            <label for="userEmail">Email</label>
                            <input type="email" class="form-control" id="userEmail" placeholder="Your Email" name="uemail">
                        </div>
                        <div class="form-group">
                            <label for="signPassword">Password</label>
                            <input type="password" class="form-control" id="signPassword" placeholder="Password" name="upasswd">
                        </div>                        
                        <button type="submit" class="btn btn-default btn-block">Login</button>
                    </form>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                </div>
            </div>
        </div>
    </div>
    <!-- signup -->
    <div class="modal fade" id="SignupModal" tabindex="-1" role="dialog" aria-labelledby="SignupLabel">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="SignupLabel">SignUp Form</h4>
                </div>
                <div class="modal-body">
                    <form action="/signup" method="POST" enctype="multipart/form-data">
                    <input type="hidden" value="put" name="_method">
                        <div class="form-group">
                            <label for="">Email</label>
                            <input type="email" class="form-control" id="" placeholder="Your Email" name="semail">
                        </div>
                        <div class="form-group">
                            <label for="userName">Name</label>
                            <input type="text" class="form-control" id="userName" placeholder="Your Name" name="sname">
                        </div>
                        <div class="form-group">
                            <label for="signPassword">Password</label>
                            <input type="password" class="form-control" id="signPassword" placeholder="Password" name="spasswd">
                        </div>
                        <div class="form-group">
                            <label for="inviteCode">inviteCode</label>
                            <input type="text" class="form-control" id="inviteCode" placeholder="Your Invite Code" name="invitecode">
                            <p class="help-block">You can ask for me !!!</p>
                        </div>
                        <div class="form-group">
                            <label for="Photograph">Photograph</label>
                            <input type="file" id="Photograph" name="photograph">
                        </div>                       
                        <button type="submit" class="btn btn-default btn-block">SignUp</button>
                    </form>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                </div>
            </div>
        </div>
    </div>

    <!--展示信息-->
    <div class="modal fade" id="shownote" tabindex="-1" role="dialog" aria-labelledby="shownote">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                
            </div>
        </div>
    </div>
    <!--模态框-->
{{end}}
